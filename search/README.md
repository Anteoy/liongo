curl -XDELETE 'http://localhost:9200/articles'

curl -XPUT 'http://127.0.0.1:9200/articles' -d '
{
	"settings": { "number_of_shards": 1 },
    "mappings" : {
        "article": {
            "properties": {
                "id": {
                    "type": "keyword",
                    "index": "false"
                },
                "title": {
                    "type": "text",
                    "term_vector": "with_positions_offsets",
                    "analyzer": "ik_max_word",
                    "search_analyzer": "ik_max_word"
                },
                "content": {
                    "type": "text",
                    "term_vector": "with_positions_offsets",
                    "analyzer": "ik_max_word",
                    "search_analyzer": "ik_max_word"
                },
                "abstract":{
                   "type": "text",
                    "term_vector": "with_positions_offsets",
                    "analyzer": "ik_max_word",
                    "search_analyzer": "ik_max_word"
                },
                "link": {
                    "type": "text"
                },
                "author": {
                    "type": "keyword",
                    "index": "true"
                },
                "classify": {
                    "type": "keyword",
                    "index" : "true"
                },
                "date": {
                    "type" : "keyword",
                    "index" : "not_analyzed"
                }
            }
        }
    }
}'



curl -XPOST "http://127.0.0.1:9200/articles/article" -H 'Content-Type: application/json' -d
'{
  "title": "芝麻分从本地mongo中获取有效数据工具",
  "date": "2017-01-16 22:00",
  "classify": "java",
  "abstract": "###引言：\n　　起因：java接入芝麻分，接口数据为了提高效率，需要储存在mongo中，若接口调用则优先从本地mongo库中查询是否存在有效数据。\n###mongo查询：\n\nimport com.alibaba.fastjson.util.TypeUtils;\nimport com.lemon.datamarket.dao.mongo.ManageRepository;\nimport com.lemon.datamarket.model.DataTransferObject;\nimport com.lemon.datamarket.utils.MongoDBUtils;\nimport com.lemon.datamarket.utils.zhima.AliConstants;\nimport com.lemon.datamarket.utils.zhima.DateUtils;\nimport com.mongodb.MongoClient;\nimport com.mongodb.client.FindIterable;\nimport com.mongodb.client.MongoCollection;\nimport com.mongodb.client.MongoCursor;\nimport org.bson.Document;\nimport org.springframework.beans.factory.annotation.Autowired;\nimport org.springframework.beans.factory.annotation.Value;\nimport org.springframework.stereotype.Component;\n\nimport javax.annotation.Resource;\nimport java.text.ParseException;\nimport java.util.Date;\nimport java.util.LinkedHashMap;\nimport java.util.Map;\n\n/**\n * Created by zhoudazhuang\n * Date: 17-1-13\n * Time: 下午3:14\n * ",
  "author": "anteoy@gmail.com",
  "link": "1/1/1/62.html",
  "content": "<p>###引言：\n　　起因：java接入芝麻分，接口数据为了提高效率，需要储存在mongo中，若接口调用则优先从本地mongo库中查询是否存在有效数据。\n###mongo查询：</p>\n\n<pre class=\"prettyprint linenums\">import com.alibaba.fastjson.util.TypeUtils;\nimport com.lemon.datamarket.dao.mongo.ManageRepository;\nimport com.lemon.datamarket.model.DataTransferObject;\nimport com.lemon.datamarket.utils.MongoDBUtils;\nimport com.lemon.datamarket.utils.zhima.AliConstants;\nimport com.lemon.datamarket.utils.zhima.DateUtils;\nimport com.mongodb.MongoClient;\nimport com.mongodb.client.FindIterable;\nimport com.mongodb.client.MongoCollection;\nimport com.mongodb.client.MongoCursor;\nimport org.bson.Document;\nimport org.springframework.beans.factory.annotation.Autowired;\nimport org.springframework.beans.factory.annotation.Value;\nimport org.springframework.stereotype.Component;\n\nimport javax.annotation.Resource;\nimport java.text.ParseException;\nimport java.util.Date;\nimport java.util.LinkedHashMap;\nimport java.util.Map;\n\n/**\n * Created by zhoudazhuang\n * Date: 17-1-13\n * Time: 下午3:14\n * Description : ali Mongo 仓库操作\n */\n@Component\npublic class AliRepository {\n\n    @Autowired\n    private MongoDBUtils mongoDBUtils;\n\n    @Value(\"${mongo.db_name}\")\n    private String databaseName;\n\n    @Resource\n    private ManageRepository manageRepository;\n\n    /**\n     * 保存集合\n     * @param collectionName\n     * @param poJo\n     */\n    public void save(String collectionName,Object poJo){\n        TypeUtils.compatibleWithJavaBean = true;\n        manageRepository.save(collectionName,poJo);\n    }\n\n    /**\n     * 芝麻信用分接口本地查询\n     * @param dataTransferObject\n     * @param type 0 手机号入参 1 姓名 身份证入参\n     * @return\n     */\n    public Map<String,Object> zhimaCreditScoreQuery(DataTransferObject dataTransferObject,int type) throws ParseException {\n        //实例化mongoClient\n        MongoClient mongoClient = mongoDBUtils.getMongoClient();\n        //获取mongo集合入口\n        MongoCollection<Document> mongoCollection = mongoClient.getDatabase(databaseName).getCollection(AliConstants.MARKET_ALI_ZHIMACREDITSCOREQUERY_MONGOCOLLECTION_NAME);\n        //芝麻信用分的数据最大有效期一个月\n        long threshold = System.currentTimeMillis() - 30 * 24 * 60 * 60 * 1000L;\n        //实例化查询对象 并组装查询参数\n        Document queryDocument = new Document();\n        if(0 == type){//要求手机号进行查询\n            queryDocument.append(\"phone\", dataTransferObject.getPhone()).append(\"createTime\", new Document().append(\"$gt\", threshold));\n        }else if(1 == type){\n            queryDocument.append(\"name\",dataTransferObject.getName()).append(\"idCard\",dataTransferObject.getIdCard()).append(\"createTime\",new Document().append(\"$gt\", threshold));\n        }\n        //获得FindIterable\n        FindIterable<Document> findIterable = mongoCollection.find(queryDocument).limit(1);//只使用第一条记录\n        MongoCursor<Document> mongoCursor = findIterable.iterator();\n        Map<String,Object> data = null;\n        //遍历FindIterable并组装对象\n        if (mongoCursor.hasNext()) {\n            //如果存在则实例化过滤规则\n            Date starttime = DateUtils.getMonthByNow(-1);//获取上个月6号 2016-12-06 00:00:00\n            Date thisdate = DateUtils.getMonthByNow(0); //获取当前月6号  2016-1-06 00:00:00\n            Date deadline = DateUtils.getMonthByNow(+1);//获取下一个6号  2016-2-06 00:00:00\n            //判断当前时间区间范围\n            Boolean bone =DateUtils.isBetweenDate(starttime,thisdate);\n            Boolean btwo =DateUtils.isBetweenDate(thisdate,deadline);\n            if(bone == true){//当前时间未超过当月6号\n                Document document = mongoCursor.next();\n                data = new LinkedHashMap<String,Object>();\n                data.put(\"result\", document.get(\"result\"));\n                data.put(\"openId\", document.get(\"open_id\"));\n            }\n            if(btwo == true){//当前时间已超过当月6号\n                Document document = mongoCursor.next();\n                //获取数据存储时间\n                Date datatime = new Date(document.getLong(\"createTime\"));\n                if(DateUtils.isBetweenDateplus(thisdate,datatime,deadline)){\n                    data = new LinkedHashMap<String,Object>();\n                    data.put(\"result\", document.get(\"result\"));\n                    data.put(\"openId\", document.get(\"open_id\"));\n                }\n            }\n        }\n        mongoCursor.close();\n\n        return data;\n    }\n\n}\n</pre>\n\n<p>###java处理时间的工具类：</p>\n\n<pre class=\"prettyprint linenums\">import java.text.ParseException;\nimport java.text.SimpleDateFormat;\nimport java.util.Calendar;\nimport java.util.Date;\n\n/**\n * Created by zhoudazhuang\n * Date: 17-1-17\n * Time: 上午10:54\n * Description :阿里时间处理\n */\npublic class DateUtils {\n    public static void main(String[] args) {\n        Date starttime = DateUtils.getMonthByNow(-1);//获取上个月6号 2016-12-06 00:00:00\n        Date thisdate = DateUtils.getMonthByNow(0); //获取当前月6号  2016-1-06 00:00:00\n        Date deadline = DateUtils.getMonthByNow(+1);//获取下一个6号  2016-2-06 00:00:00\n        Boolean b1 = DateUtils.isBetweenDate(starttime,thisdate);\n        Boolean b2 = DateUtils.isBetweenDate(thisdate,deadline);\n        System.out.println(b1);\n        System.out.println(b2);\n    }\n\n    /**\n     * 对比当前时间是否在指定区间内\n     * @param starttime  starttime < now time < deadline\n     * @param deadline\n     * @return\n     */\n    public static boolean isBetweenDate(Date starttime,Date deadline){\n        Date date = new Date();//需要修改这里为mongo拿出的芝麻数据存入时间\n        if(date.after(starttime) && date.before(deadline)){\n            return true;\n        }\n        return false;\n    }\n\n    /**\n     * 对比第二个时间是否在指定第一个和最后一个区间内\n     * @param starttime  starttime < now time < deadline\n     * @param deadline\n     * @return\n     */\n    public static boolean isBetweenDateplus(Date starttime,Date standertime,Date deadline){\n        if(standertime.after(starttime) && standertime.before(deadline)){\n            return true;\n        }\n        return false;\n    }\n\n    /**\n     * 获取指定月的6号\n     * @param month -1 从当前月减一月 +1 从当前月加一月\n     * @return\n     */\n    public static Date getMonthByNow(int month){\n        SimpleDateFormat simpleDateFormatMM = new SimpleDateFormat(\"yyyy-MM\");\n        SimpleDateFormat simpleDateFormat = new SimpleDateFormat(\"yyyy-MM-dd HH:mm:ss\");\n        Calendar c = Calendar.getInstance();\n\n        c.setTime(new Date());\n        c.add(Calendar.MONTH, month);\n        Date mb = c.getTime();\n        String mon = simpleDateFormatMM.format(mb);\n\n        StringBuffer stringBuffer = new StringBuffer();\n        stringBuffer.append(mon).append(\"-06 0:0:0\");\n\n        System.out.println(stringBuffer.toString());\n        Date date = null;\n        try {\n            date = simpleDateFormat.parse(stringBuffer.toString());\n        } catch (ParseException e) {\n            e.printStackTrace();\n        }\n        return date;\n    }\n\n\n}\n\n</pre>\n\n<p>###附录\n\t这里记录下mysql的时间以及order by，包括int,varchar,timestamp,data\n表：</p>\n\n<pre class=\"prettyprint linenums\">DROP TABLE IF EXISTS `test`;\n/*!40101 SET @saved_cs_client     = @@character_set_client */;\n/*!40101 SET character_set_client = utf8 */;\nCREATE TABLE `test` (\n  `id` int(11) NOT NULL,\n  `time` varchar(45) DEFAULT NULL,\n  `timestamp` timestamp NULL DEFAULT NULL,\n  `date` date DEFAULT NULL,\n  `str` varchar(45) DEFAULT NULL,\n  PRIMARY KEY (`id`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n</pre>\n\n<p>测试sql</p>\n\n<pre class=\"prettyprint linenums\">SELECT * FROM chargeManager.test;\nset sql_safe_updates = 0;\nselect unix_timestamp('2009-10-26 10-06-07');\nUPDATE `chargeManager`.`test` SET `timestamp`=1256522767 WHERE `id`='1';\nSELECT * FROM chargeManager.test where str >2 order by id desc;#int\nSELECT * FROM chargeManager.test where date > '2009-10-26' order by time desc;#varchar date > 需要 ‘’\nSELECT * FROM chargeManager.test where time > 100 order by test.timestamp desc;#timestamp\nSELECT * FROM chargeManager.test where id>1 order by test.date desc;\n</pre>\n\n<p><em>注：varchar中若为”123”等数字类型字符串，则类int排序，而timestamp插入时间戳，由于在时间戳的存储中常用int或varchar，所以注意若使用sql操作timestamp时，传入的需为datatime类型，既不是int，也不是varchar.</em></p>\n",
  "id": "62"
}'