---
date: 2017-01-16 22:00
title: 芝麻分从本地mongo中获取有效数据工具
categories:
    - java
tags:
    - java
---

###引言：
　　起因：java接入芝麻分，接口数据为了提高效率，需要储存在mongo中，若接口调用则优先从本地mongo库中查询是否存在有效数据。
###mongo查询：

```
import com.alibaba.fastjson.util.TypeUtils;
import com.lemon.datamarket.dao.mongo.ManageRepository;
import com.lemon.datamarket.model.DataTransferObject;
import com.lemon.datamarket.utils.MongoDBUtils;
import com.lemon.datamarket.utils.zhima.AliConstants;
import com.lemon.datamarket.utils.zhima.DateUtils;
import com.mongodb.MongoClient;
import com.mongodb.client.FindIterable;
import com.mongodb.client.MongoCollection;
import com.mongodb.client.MongoCursor;
import org.bson.Document;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import javax.annotation.Resource;
import java.text.ParseException;
import java.util.Date;
import java.util.LinkedHashMap;
import java.util.Map;

/**
 * Created by zhoudazhuang
 * Date: 17-1-13
 * Time: 下午3:14
 * Description : ali Mongo 仓库操作
 */
@Component
public class AliRepository {

    @Autowired
    private MongoDBUtils mongoDBUtils;

    @Value("${mongo.db_name}")
    private String databaseName;

    @Resource
    private ManageRepository manageRepository;

    /**
     * 保存集合
     * @param collectionName
     * @param poJo
     */
    public void save(String collectionName,Object poJo){
        TypeUtils.compatibleWithJavaBean = true;
        manageRepository.save(collectionName,poJo);
    }

    /**
     * 芝麻信用分接口本地查询
     * @param dataTransferObject
     * @param type 0 手机号入参 1 姓名 身份证入参
     * @return
     */
    public Map<String,Object> zhimaCreditScoreQuery(DataTransferObject dataTransferObject,int type) throws ParseException {
        //实例化mongoClient
        MongoClient mongoClient = mongoDBUtils.getMongoClient();
        //获取mongo集合入口
        MongoCollection<Document> mongoCollection = mongoClient.getDatabase(databaseName).getCollection(AliConstants.MARKET_ALI_ZHIMACREDITSCOREQUERY_MONGOCOLLECTION_NAME);
        //芝麻信用分的数据最大有效期一个月
        long threshold = System.currentTimeMillis() - 30 * 24 * 60 * 60 * 1000L;
        //实例化查询对象 并组装查询参数
        Document queryDocument = new Document();
        if(0 == type){//要求手机号进行查询
            queryDocument.append("phone", dataTransferObject.getPhone()).append("createTime", new Document().append("$gt", threshold));
        }else if(1 == type){
            queryDocument.append("name",dataTransferObject.getName()).append("idCard",dataTransferObject.getIdCard()).append("createTime",new Document().append("$gt", threshold));
        }
        //获得FindIterable
        FindIterable<Document> findIterable = mongoCollection.find(queryDocument).limit(1);//只使用第一条记录
        MongoCursor<Document> mongoCursor = findIterable.iterator();
        Map<String,Object> data = null;
        //遍历FindIterable并组装对象
        if (mongoCursor.hasNext()) {
            //如果存在则实例化过滤规则
            Date starttime = DateUtils.getMonthByNow(-1);//获取上个月6号 2016-12-06 00:00:00
            Date thisdate = DateUtils.getMonthByNow(0); //获取当前月6号  2016-1-06 00:00:00
            Date deadline = DateUtils.getMonthByNow(+1);//获取下一个6号  2016-2-06 00:00:00
            //判断当前时间区间范围
            Boolean bone =DateUtils.isBetweenDate(starttime,thisdate);
            Boolean btwo =DateUtils.isBetweenDate(thisdate,deadline);
            if(bone == true){//当前时间未超过当月6号
                Document document = mongoCursor.next();
                data = new LinkedHashMap<String,Object>();
                data.put("result", document.get("result"));
                data.put("openId", document.get("open_id"));
            }
            if(btwo == true){//当前时间已超过当月6号
                Document document = mongoCursor.next();
                //获取数据存储时间
                Date datatime = new Date(document.getLong("createTime"));
                if(DateUtils.isBetweenDateplus(thisdate,datatime,deadline)){
                    data = new LinkedHashMap<String,Object>();
                    data.put("result", document.get("result"));
                    data.put("openId", document.get("open_id"));
                }
            }
        }
        mongoCursor.close();

        return data;
    }

}
```

###java处理时间的工具类：
```
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Calendar;
import java.util.Date;

/**
 * Created by zhoudazhuang
 * Date: 17-1-17
 * Time: 上午10:54
 * Description :阿里时间处理
 */
public class DateUtils {
    public static void main(String[] args) {
        Date starttime = DateUtils.getMonthByNow(-1);//获取上个月6号 2016-12-06 00:00:00
        Date thisdate = DateUtils.getMonthByNow(0); //获取当前月6号  2016-1-06 00:00:00
        Date deadline = DateUtils.getMonthByNow(+1);//获取下一个6号  2016-2-06 00:00:00
        Boolean b1 = DateUtils.isBetweenDate(starttime,thisdate);
        Boolean b2 = DateUtils.isBetweenDate(thisdate,deadline);
        System.out.println(b1);
        System.out.println(b2);
    }

    /**
     * 对比当前时间是否在指定区间内
     * @param starttime  starttime < now time < deadline
     * @param deadline
     * @return
     */
    public static boolean isBetweenDate(Date starttime,Date deadline){
        Date date = new Date();//需要修改这里为mongo拿出的芝麻数据存入时间
        if(date.after(starttime) && date.before(deadline)){
            return true;
        }
        return false;
    }

    /**
     * 对比第二个时间是否在指定第一个和最后一个区间内
     * @param starttime  starttime < now time < deadline
     * @param deadline
     * @return
     */
    public static boolean isBetweenDateplus(Date starttime,Date standertime,Date deadline){
        if(standertime.after(starttime) && standertime.before(deadline)){
            return true;
        }
        return false;
    }

    /**
     * 获取指定月的6号
     * @param month -1 从当前月减一月 +1 从当前月加一月
     * @return
     */
    public static Date getMonthByNow(int month){
        SimpleDateFormat simpleDateFormatMM = new SimpleDateFormat("yyyy-MM");
        SimpleDateFormat simpleDateFormat = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        Calendar c = Calendar.getInstance();

        c.setTime(new Date());
        c.add(Calendar.MONTH, month);
        Date mb = c.getTime();
        String mon = simpleDateFormatMM.format(mb);

        StringBuffer stringBuffer = new StringBuffer();
        stringBuffer.append(mon).append("-06 0:0:0");

        System.out.println(stringBuffer.toString());
        Date date = null;
        try {
            date = simpleDateFormat.parse(stringBuffer.toString());
        } catch (ParseException e) {
            e.printStackTrace();
        }
        return date;
    }


}

```
###附录
	这里记录下mysql的时间以及order by，包括int,varchar,timestamp,data
表：

```
DROP TABLE IF EXISTS `test`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `test` (
  `id` int(11) NOT NULL,
  `time` varchar(45) DEFAULT NULL,
  `timestamp` timestamp NULL DEFAULT NULL,
  `date` date DEFAULT NULL,
  `str` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```
测试sql

```
SELECT * FROM chargeManager.test;
set sql_safe_updates = 0;
select unix_timestamp('2009-10-26 10-06-07');
UPDATE `chargeManager`.`test` SET `timestamp`=1256522767 WHERE `id`='1';
SELECT * FROM chargeManager.test where str >2 order by id desc;#int
SELECT * FROM chargeManager.test where date > '2009-10-26' order by time desc;#varchar date > 需要 ‘’
SELECT * FROM chargeManager.test where time > 100 order by test.timestamp desc;#timestamp
SELECT * FROM chargeManager.test where id>1 order by test.date desc;
```
*注：varchar中若为"123"等数字类型字符串，则类int排序，而timestamp插入时间戳，由于在时间戳的存储中常用int或varchar，所以注意若使用sql操作timestamp时，传入的需为datatime类型，既不是int，也不是varchar.*
