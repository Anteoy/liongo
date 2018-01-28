---
date: 2016-10-03 00:05:00
title: 关于MD5，SALT与SHA1的部分HASH算法解析
categories:
    - 加密
tags:
---

   在这里我只贴下简单实现，有兴趣了解shiro salt的，大家可以去看看源码哈，或者可以自己实现一些简单的md5，sha1,base64等的简易salt实现，以下示例基本都有注释，如有不对，欢迎指正

```
package com.zy.md;

import java.security.MessageDigest;
import java.util.Random;

/**
 * Created by root on 16-10-2.
 */
public class Test {
    //main测试类
    public static void main(String[] args) {
        String result = getMD5("zhoudazhuang");
        System.err.println(result);
    }

    /**
     * 生成md5
     * @param message
     * @return
     */
    public static String getMD5(String message) {
        String md5str = "";
        try {
            String token = System.currentTimeMillis()+new Random().nextInt()+"";
            System.out.println(token);
            //1 创建一个提供信息摘要算法的对象，初始化为md5算法对象
            MessageDigest md = MessageDigest.getInstance("MD5");//这里可换为SHA1
            //2 将消息变成byte数组
            byte[] input = message.getBytes();
            //3 计算后获得字节数组,这就是那128位了 byte表示为-128 127 MD5最终需32位的16进制字符表示
            byte[] buff = md.digest(input);
            System.out.printf(buff.toString());
            //4 把数组每一字节（一个字节占八位）换成16进制连成md5字符串
            md5str = bytesToHex(buff);

        } catch (Exception e) {
            e.printStackTrace();
        }
        return md5str;
    }

    /**
     * Integer.toHexString取得了十六进制字符串,可以看出
     * b[i] & 0xFF运算后得出的仍然是个int,那么为何要和 0xFF进行与运算呢?直接 Integer.toHexString(b[i]);将byte强转为int不行吗?答案是不行的.
     * 其原因在于:
     * 1.byte的大小为8bits而int的大小为32bits
     * 2.java的二进制采用的是补码形式
     * 二进制转十六进制
     * @param bytes
     * @return
     */
    public static String bytesToHex(byte[] bytes) {
        StringBuffer md5str = new StringBuffer();
        //把数组每一字节换成16进制连成md5字符串
        int digital;
        for (int i = 0; i < bytes.length; i++) {
            digital = bytes[i];
            //Java中的一个byte，其范围是-128~127的，而Integer.toHexString的参数本来是int，如果不进行&0xff，那么当一个byte会转换成int时，对于负数，会做位扩展，举例来说，一个byte的-1（即0xff），会被转换成int的-1（即0xffffffff），那么转化出的结果就不是我们想要的了。
            //而0xff默认是整形，所以，一个byte跟0xff相与会先将那个byte转化成整形运算，这样，结果中的高的24个比特就总会被清0，于是结果总是我们想要的
            //而0xff为255 所以这里需使用 digital += 256
            if(digital < 0) {
                digital += 256;
            }
            //当digital比16小时,高位补0，否则位数不够
            if(digital < 16){
                md5str.append("0");
            }
            md5str.append(Integer.toHexString(digital));
        }
        return md5str.toString().toUpperCase();
    }


    public static String byteArrayToHex(byte[] byteArray) {

        // 首先初始化一个字符数组，用来存放每个16进制字符
        char[] hexDigits = {'0','1','2','3','4','5','6','7','8','9', 'A','B','C','D','E','F' };
        // new一个字符数组，这个就是用来组成结果字符串的（解释一下：一个byte是八位二进制，也就是2位十六进制字符（2的8次方等于16的2次方））
        char[] resultCharArray =new char[byteArray.length * 2];
        // 遍历字节数组，通过位运算（位运算效率高），转换成字符放到字符数组中去
        int index = 0;
        for (byte b : byteArray) {

            //hexDigits[b>>> 4 & 0xf] 由b获得16进制中的高8位 无符号右移4位与0xf 无符号右移4位则是计算高位
            resultCharArray[index++] = hexDigits[b>>> 4 & 0xf];
            //hexDigits[b>>> 4 & 0xf] 由b获得16进制中的低8位 b与0xf
            resultCharArray[index++] = hexDigits[b& 0xf];

        }
        // 字符数组组合成字符串返回
        return new String(resultCharArray);
    }
}

```


