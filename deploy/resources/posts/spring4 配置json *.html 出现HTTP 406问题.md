---
date: 2016-11-10 17:35:00
title: spring4 配置json *.html 出现HTTP 406问题
categories:
    - java
tags:
    - spring
---


**引起原因：**


 1. spring4.x中原请求servlet-mapping为/ 拦截所有请求

 2. 需要修改为*.html,需使用spring4.x

 3. 重现： 原本是直接修改servlet-mapping  /改位*.html

 4. 结果：修改为*.html过后 请求无法正常发挥Json,一直报错htpp 406 无法接受的请求头


**寻找问题并解决过程：**


 - 发现spring-servlet.xml的xsi原始指向位spring3.x 果断修改为4.x 后来干脆直接删掉后缀，如下：


```
<beans xmlns="http://www.springframework.org/schema/beans"
	xmlns:context="http://www.springframework.org/schema/context" xmlns:p="http://www.springframework.org/schema/p"
	xmlns:mvc="http://www.springframework.org/schema/mvc" xmlns:util="http://www.springframework.org/schema/util"
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:tx="http://www.springframework.org/schema/tx"
	xsi:schemaLocation="http://www.springframework.org/schema/beans
      http://www.springframework.org/schema/beans/spring-beans.xsd
      http://www.springframework.org/schema/context
      http://www.springframework.org/schema/context/spring-context.xsd
      http://www.springframework.org/schema/mvc
      http://www.springframework.org/schema/mvc/spring-mvc.xsd
      http://www.springframework.org/schema/util
      http://www.springframework.org/schema/util/spring-util.xsd
      http://www.springframework.org/schema/tx
      http://www.springframework.org/schema/tx/spring-tx.xsd">
```
    仍然不能解决。

 - google,baidu一堆，发现spring3.x和spring4.x所使用的jar包依赖，以及配置不尽相同，最后尝试了很多版本 比如：
 1.*stackoverflow*


```
<!-- activates annotation driven binding -->
<mvc:annotation-driven ignoreDefaultModelOnRedirect="true" validator="validator">
    <mvc:message-converters>
        <bean class="org.springframework.http.converter.ResourceHttpMessageConverter"/>
        <bean class="org.springframework.http.converter.xml.Jaxb2RootElementHttpMessageConverter"/>
        <bean class="org.springframework.http.converter.json.MappingJacksonHttpMessageConverter"/>
    </mvc:message-converters>
</mvc:annotation-driven>
```

```
@RequestMapping(value = "/your_url", method = RequestMethod.GET, produces = "application/json")
@ResponseBody
```



   2.*ifo*
```
<mvc:annotation-driven>
		<mvc:message-converters register-defaults="true">
			<!-- 避免IE执行AJAX时,返回JSON出现下载文件 -->
			<bean id="fastJsonHttpMessageConverter" class="com.alibaba.fastjson.support.spring.FastJsonHttpMessageConverter">
				<property name="supportedMediaTypes">
					<list>
						<value>application/json;charset=UTF-8</value>
					</list>
				</property>
			</bean>
		</mvc:message-converters>
	</mvc:annotation-driven>
```
3.csdn
```
	<!-- 这里处理 json 的转换 -->
	<bean id="mappingJacksonHttpMessageConverter"
		class="org.springframework.http.converter.json.MappingJacksonHttpMessageConverter">
		<property name="supportedMediaTypes">
			<list>
				<value>application/json;charset=UTF-8</value>
				<value>text/html;charset=UTF-8</value>
			</list>
		</property>
	</bean>
```
	配合
```<dependency>
            <groupId>org.codehaus.jackson</groupId>
            <artifactId>jackson-mapper-asl</artifactId>
            <version>1.9.13</version>
        </dependency>
```
        **这里的配置在spring3.x是可用的，按理在spring4.x为弃用，但后来同事证明他目前的dubbo使用的就是此种配置，后来我试了试，虽然使用/配置servlet-mapping可用，但一旦配置为*.html，便会同样出现406错误**

**解决问题：**
  经过一番折腾，发现
  1. spring4.x 在处理Json的时候 同spring3.x并不相同
  2. spring4.x 依赖jar以及配置都需要作调整



----------


  *结果如下*


 - pom.xml
```1.spring3.x```

```
	 <dependency>
            <groupId>org.codehaus.jackson</groupId>
            <artifactId>jackson-mapper-asl</artifactId>
            <version>1.9.13</version>
        </dependency>
```

      2.spring4.x

```
<dependency>
            <groupId>com.fasterxml.jackson.core</groupId>
            <artifactId>jackson-databind</artifactId>
            <version>2.6.3</version>
        </dependency>
        <dependency>
            <groupId>com.fasterxml.jackson.core</groupId>
            <artifactId>jackson-core</artifactId>
            <version>2.6.3</version>
        </dependency>
        <dependency>
            <groupId>com.fasterxml.jackson.core</groupId>
            <artifactId>jackson-annotations</artifactId>
            <version>2.6.3</version>
        </dependency>
```

 - spring-servlet.xml
 ```1.spring3.x```

```<!-- 这里处理 json 的转换 -->
	<bean id="mappingJacksonHttpMessageConverter"
		class="org.springframework.http.converter.json.MappingJacksonHttpMessageConverter">
		<property name="supportedMediaTypes">
			<list>
				<value>application/json;charset=UTF-8</value>
				<value>text/html;charset=UTF-8</value>
			</list>
		</property>
	</bean>
```
*注：这个在spring4.x中也可以使用，但不能修改servlet-mapping为*.html 否则会报错http 406
                   ```  2. spring4.x```
```
<!-- 启动注解驱动的Spring MVC功能，注册请求url和注解POJO类方法的映射 -->
	<mvc:annotation-driven content-negotiation-manager="contentNegotiationManager"/>
	<bean id="contentNegotiationManager" class="org.springframework.web.accept.ContentNegotiationManagerFactoryBean">
		<property name="favorPathExtension" value="false" />
	</bean>
	<bean id="jsonConverter" class="org.springframework.http.converter.json.MappingJackson2HttpMessageConverter"></bean>

	<bean id="stringConverter"
		  class="org.springframework.http.converter.StringHttpMessageConverter">
		<property name="supportedMediaTypes">
			<list>
				<value>text/plain;charset=UTF-8</value>
			</list>
		</property>
	</bean>
```