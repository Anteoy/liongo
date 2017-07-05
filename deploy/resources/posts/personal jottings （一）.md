---
date: 2016-10-26 14:50:00
title: personal jottings （一）
categories:
    - jottings
tags:
    - jottings
---

# 个人杂记（一）

```
 <!--<build>
        <plugins>
            <plugin>
                <groupId>org.apache.maven.plugins</groupId>
                <artifactId>maven-shade-plugin</artifactId>
                <executions>
                    <execution>
                        <phase>package</phase>
                        <goals>
                            <goal>shade</goal>
                        </goals>
                        <configuration>
                            <transformers>
                                <transformer
                                        implementation="org.apache.maven.plugins.shade.resource.AppendingTransformer">
                                    <resource>META-INF/spring.handlers</resource>
                                </transformer>
                                <transformer
                                        implementation="org.apache.maven.plugins.shade.resource.AppendingTransformer">
                                    <resource>META-INF/spring.schemas</resource>
                                </transformer>
                                <transformer
                                        implementation="org.apache.maven.plugins.shade.resource.ManifestResourceTransformer">
                                    <mainClass>com.lemon.mainservice.ApplyTask</mainClass>
                                </transformer>
                                <transformer
                                        implementation="org.apache.maven.plugins.shade.resource.ServicesResourceTransformer" />

                            </transformers>
                        </configuration>
                    </execution>
                </executions>
            </plugin>
            <plugin>
                <groupId>org.apache.maven.plugins</groupId>
                <artifactId>maven-compiler-plugin</artifactId>
                <configuration>
                    <source>1.7</source>
                    <target>1.7</target>
                </configuration>
            </plugin>
        </plugins>
        <resources>
            <resource>
                <directory>${basedir}/src/main/resources</directory>
                <includes>
                    <include>**/*.properties</include>
                    <include>**/*.xml</include>
                </includes>
            </resource>
            <resource>
                <directory>${basedir}/src/main/java</directory>
                <includes>
                    <include>**/*.xml</include>
                    <include>**/*.properties</include>
                </includes>
            </resource>
        </resources>
    </build>-->
    <!--<build>
        <pluginManagement>
            <plugins>
                <plugin>
                    <groupId>org.apache.maven.plugins</groupId>
                    <artifactId>maven-compiler-plugin</artifactId>
                    <configuration>
                        <source>1.7</source>
                        <target>1.7</target>
                    </configuration>
                </plugin>
                <plugin>
                    <groupId>org.apache.maven.plugins</groupId>
                    <artifactId>maven-jar-plugin</artifactId>
                    <configuration>
                        <archive>
                            <manifest>
                                <mainClass>com.lemon.mainservice.ApplyTask</mainClass>
                                <addClasspath>true</addClasspath>
                                <classpathPrefix>lib/</classpathPrefix>
                            </manifest>

                        </archive>
                        <classesDirectory>
                        </classesDirectory>
                    </configuration>
                </plugin>
            </plugins>
        </pluginManagement>
    </build> TODO lib无法打入jar -->
    <build>
		<plugins>
			<plugin>
				<artifactId>maven-assembly-plugin</artifactId>
				<configuration>
					<archive>
						<manifest>
							<mainClass>com.lemon.mainservice.ApplyTask</mainClass>
						</manifest>
					</archive>
					<descriptorRefs>
						<descriptorRef>jar-with-dependencies</descriptorRef>
					</descriptorRefs>
				</configuration>
			</plugin>
		</plugins>
	</build>
```
执行：

```
ps -ef |grep jar
ps -aux |grep jar(两者风格略有差异)
jps
java -jar batchExport-0.0.1-SNAPSHOT-jar-with-dependencies.jar &
history
nohup java -cp batchExport-0.0.1-SNAPSHOT-jar-with-dependencies.jar -Xms1024M -Xmx2048M -XX:PermSize=100M -XX:MaxPermSize=100M com.lemon.mainservice.ApplyTask &(不挂断地运行命令)


```

```
sudo ufw status
一个相对iptables简单很多的防火墙配置工具 ubuntu fire wall
基于iptable之上
```

java泛型好处
1.消除强制类型转换
2.类型安全
3.代码重用
4.潜在的性能收益
如：
Map m = new HashMap();
m.put("key", "blarg");
String s = (String) m.get("key");

```
   public static void main(String[] args) {
        Long a = Long.valueOf(5);
        String b = "abc";
        Map c = new HashMap();
        c.put("a","ooo");
        oo(c);
        System.out.println(c);
    }

    public static void  oo(Long a){
        a = Long.valueOf(4);
    }
    public static void  oo(String a){
        a = a+"abc";
    }
    public static void  oo(Map a){
        a.put("a","uuu");
    }
}
只有map等对象才会引用传递，包装类和基本类型 都是值传递
```

override（重写，覆盖）
1、方法名、参数、返回值相同。
2、子类方法不能缩小父类方法的访问权限。
3、子类方法不能抛出比父类方法更多的异常(但子类方法可以不抛出异常)。
4、存在于父类和子类之间。
5、方法被定义为final不能被重写。

overload（重载，过载）
1、参数类型、个数、顺序至少有一个不相同。
2、不能重载只有返回值不同的方法名。
3、存在于父类和子类、同类中。

sql注入方法
方法1
先猜表名
And (Select count(*) from 表名)<>0
猜列名
And (Select count（列名） from 表名）<>0
或者也可以这样
and exists (select * from 表名）
and exists (select 列名 from 表名）
返回正确的，那么写的表名或列名就是正确
这里要注意的是，exists这个不能应用于猜内容上，例如and exists (select len(user) from admin)>3 这样是不行的
很多人都是喜欢查询里面的内容，一旦iis没有关闭错误提示的，那么就可以利用报错方法轻松获得库里面的内容
获得数据库连接用户名：；and user>0


1.1 什么是XSS攻击
XSS是一种经常出现在web应用中的计算机安全漏洞，它允许恶意web用户将代码植入到提供给其它用户使用的页面中。比如这些代码包括HTML代码和客户端脚本。攻击者利用XSS漏洞旁路掉访问控制——例如同源策略(same origin policy)。这种类型的漏洞由于被黑客用来编写危害性更大的网络钓鱼(Phishing)攻击而变得广为人知。对于跨站脚本攻击，黑客界共识是：跨站脚本攻击是新型的“缓冲区溢出攻击“，而JavaScript是新型的“ShellCode”。
数据来源：2007 OWASP Top 10的MITRE数据
注：OWASP是世界上最知名的Web安全与数据库安全研究组织
在2007年OWASP所统计的所有安全威胁中，跨站脚本攻击占到了22%，高居所有Web威胁之首。
XSS攻击的危害包括
1、盗取各类用户帐号，如机器登录帐号、用户网银帐号、各类管理员帐号
2、控制企业数据，包括读取、篡改、添加、删除企业敏感数据的能力
3、盗窃企业重要的具有商业价值的资料
4、非法转账
5、强制发送电子邮件
6、网站挂马
7、控制受害者机器向其它网站发起攻击
1.2 XSS漏洞的分类
XSS漏洞按照攻击利用手法的不同，有以下三种类型：
类型A，本地利用漏洞，这种漏洞存在于页面中客户端脚本自身。其攻击过程如下所示：
Alice给Bob发送一个恶意构造了Web的URL。
Bob点击并查看了这个URL。
恶意页面中的JavaScript打开一个具有漏洞的HTML页面并将其安装在Bob电脑上。
具有漏洞的HTML页面包含了在Bob电脑本地域执行的JavaScript。
Alice的恶意脚本可以在Bob的电脑上执行Bob所持有的权限下的命令。
类型B，反射式漏洞，这种漏洞和类型A有些类似，不同的是Web客户端使用Server端脚本生成页面为用户提供数据时，如果未经验证的用户数据被包含在页面中而未经HTML实体编码，客户端代码便能够注入到动态页面中。其攻击过程如下：
Alice经常浏览某个网站，此网站为Bob所拥有。Bob的站点运行Alice使用用户名/密码进行登录，并存储敏感信息(比如银行帐户信息)。
Charly发现Bob的站点包含反射性的XSS漏洞。
Charly编写一个利用漏洞的URL，并将其冒充为来自Bob的邮件发送给Alice。
Alice在登录到Bob的站点后，浏览Charly提供的URL。
嵌入到URL中的恶意脚本在Alice的浏览器中执行，就像它直接来自Bob的服务器一样。此脚本盗窃敏感信息(授权、信用卡、帐号信息等)然后在Alice完全不知情的情况下将这些信息发送到Charly的Web站点。
类型C，存储式漏洞，该类型是应用最为广泛而且有可能影响到Web服务器自身安全的漏洞，骇客将攻击脚本上传到Web服务器上，使得所有访问该页面的用户都面临信息泄漏的可能，其中也包括了Web服务器的管理员。其攻击过程如下：
Bob拥有一个Web站点，该站点允许用户发布信息/浏览已发布的信息。
Charly注意到Bob的站点具有类型C的XSS漏洞。
Charly发布一个热点信息，吸引其它用户纷纷阅读。
Bob或者是任何的其他人如Alice浏览该信息，其会话cookies或者其它信息将被Charly盗走。
类型A直接威胁用户个体，而类型B和类型C所威胁的对象都是企业级Web应用。
