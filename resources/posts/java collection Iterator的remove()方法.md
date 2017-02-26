---
date: 2016-09-12 17:09:00
title: java collection Iterator的remove()方法
categories:
    - java
tags:
---

　　事情是这样的，今天在项目里用spring data jpa hibernate的SearchFilter返回一个Ｌist<Ｏbject>的时候，由于需求需要，必须再使用一定条件过滤部分元素，开始想iterator.remove()方法三下五除二就能解决，结果就悲剧了,前前后后搞了近一小时。下面上代码，同时记录一下：
　　

```
Specification<Lawyer> specification = DynamicSpecifications.bySearchFilter(Lawyer.class, set);
List<Lawyer> lawyerList =lawyerService.findByExample(specification, page);
```
一开始使用

```
Iterator<Lawyer> iterator = lawyerList.iterator();
        if("1".equals(customer.getType())){
            while (iterator.hasNext()){
                Lawyer lawyer = iterator.next();
                if(lawyer.getLawyerid().equals(customer.getCustomerid())){
                    iterator.remove();
                }
            }
        }
```
一直报错：

```
16:54:31.469 [http-bio-8080-exec-2] ERROR 500.jsp - java.lang.UnsupportedOperationException
	at java.util.Collections$UnmodifiableCollection$1.remove(Collections.java:1069)
	at com.lawcall.controller.CircleUserController.getlawyer(CircleUserController.java:364)
	at sun.reflect.NativeMethodAccessorImpl.invoke0(Native Method)
	at sun.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:57)
	at sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)
	at java.lang.reflect.Method.invoke(Method.java:606)
	at org.springframework.web.method.support.InvocableHandlerMethod.doInvoke(InvocableHandlerMethod.java:222)
	at org.springframework.web.method.support.InvocableHandlerMethod.invokeForRequest(InvocableHandlerMethod.java:137)
	at org.springframework.web.servlet.mvc.method.annotation.ServletInvocableHandlerMethod.invokeAndHandle(ServletInvocableHandlerMethod.java:110)
	at org.springframework.web.servlet.mvc.method.annotation.RequestMappingHandlerAdapter.invokeHandlerMethod(RequestMappingHandlerAdapter.java:814)
	at org.springframework.web.servlet.mvc.method.annotation.RequestMappingHandlerAdapter.handleInternal(RequestMappingHandlerAdapter.java:737)
	at org.springframework.web.servlet.mvc.method.AbstractHandlerMethodAdapter.handle(AbstractHandlerMethodAdapter.java:85)
	at org.springframework.web.servlet.DispatcherServlet.doDispatch(DispatcherServlet.java:959)
	at org.springframework.web.servlet.DispatcherServlet.doService(DispatcherServlet.java:893)
	at org.springframework.web.servlet.FrameworkServlet.processRequest(FrameworkServlet.java:969)
	at org.springframework.web.servlet.FrameworkServlet.doPost(FrameworkServlet.java:871)
	at javax.servlet.http.HttpServlet.service(HttpServlet.java:650)
	at org.springframework.web.servlet.FrameworkServlet.service(FrameworkServlet.java:845)
	at javax.servlet.http.HttpServlet.service(HttpServlet.java:731)
	at org.apache.catalina.core.ApplicationFilterChain.internalDoFilter(ApplicationFilterChain.java:303)
	at org.apache.catalina.core.ApplicationFilterChain.doFilter(ApplicationFilterChain.java:208)
	at org.apache.tomcat.websocket.server.WsFilter.doFilter(WsFilter.java:52)
	at org.apache.catalina.core.ApplicationFilterChain.internalDoFilter(ApplicationFilterChain.java:241)
	at org.apache.catalina.core.ApplicationFilterChain.doFilter(ApplicationFilterChain.java:208)
	at org.apache.shiro.web.servlet.ProxiedFilterChain.doFilter(ProxiedFilterChain.java:61)
	at org.apache.shiro.web.servlet.AdviceFilter.executeChain(AdviceFilter.java:108)
	at org.apache.shiro.web.servlet.AdviceFilter.doFilterInternal(AdviceFilter.java:137)
	at org.apache.shiro.web.servlet.OncePerRequestFilter.doFilter(OncePerRequestFilter.java:125)
	at org.apache.shiro.web.servlet.ProxiedFilterChain.doFilter(ProxiedFilterChain.java:66)
	at org.apache.shiro.web.servlet.AbstractShiroFilter.executeChain(AbstractShiroFilter.java:449)
	at org.apache.shiro.web.servlet.AbstractShiroFilter$1.call(AbstractShiroFilter.java:365)
	at org.apache.shiro.subject.support.SubjectCallable.doCall(SubjectCallable.java:90)
	at org.apache.shiro.subject.support.SubjectCallable.call(SubjectCallable.java:83)
	at org.apache.shiro.subject.support.DelegatingSubject.execute(DelegatingSubject.java:383)
	at org.apache.shiro.web.servlet.AbstractShiroFilter.doFilterInternal(AbstractShiroFilter.java:362)
	at org.apache.shiro.web.servlet.OncePerRequestFilter.doFilter(OncePerRequestFilter.java:125)
	at org.springframework.web.filter.DelegatingFilterProxy.invokeDelegate(DelegatingFilterProxy.java:346)
	at org.springframework.web.filter.DelegatingFilterProxy.doFilter(DelegatingFilterProxy.java:262)
	at org.apache.catalina.core.ApplicationFilterChain.internalDoFilter(ApplicationFilterChain.java:241)
	at org.apache.catalina.core.ApplicationFilterChain.doFilter(ApplicationFilterChain.java:208)
	at org.springframework.orm.jpa.support.OpenEntityManagerInViewFilter.doFilterInternal(OpenEntityManagerInViewFilter.java:178)
	at org.springframework.web.filter.OncePerRequestFilter.doFilter(OncePerRequestFilter.java:107)
	at org.apache.catalina.core.ApplicationFilterChain.internalDoFilter(ApplicationFilterChain.java:241)
	at org.apache.catalina.core.ApplicationFilterChain.doFilter(ApplicationFilterChain.java:208)
	at org.springframework.web.filter.CharacterEncodingFilter.doFilterInternal(CharacterEncodingFilter.java:121)
	at org.springframework.web.filter.OncePerRequestFilter.doFilter(OncePerRequestFilter.java:107)
	at org.apache.catalina.core.ApplicationFilterChain.internalDoFilter(ApplicationFilterChain.java:241)
	at org.apache.catalina.core.ApplicationFilterChain.doFilter(ApplicationFilterChain.java:208)
	at org.apache.catalina.core.StandardWrapperValve.invoke(StandardWrapperValve.java:218)
	at org.apache.catalina.core.StandardContextValve.invoke(StandardContextValve.java:122)
	at org.apache.catalina.authenticator.AuthenticatorBase.invoke(AuthenticatorBase.java:505)
	at org.apache.catalina.core.StandardHostValve.invoke(StandardHostValve.java:169)
	at org.apache.catalina.valves.ErrorReportValve.invoke(ErrorReportValve.java:103)
	at org.apache.catalina.valves.AccessLogValve.invoke(AccessLogValve.java:956)
	at org.apache.catalina.core.StandardEngineValve.invoke(StandardEngineValve.java:116)
	at org.apache.catalina.connector.CoyoteAdapter.service(CoyoteAdapter.java:442)
	at org.apache.coyote.http11.AbstractHttp11Processor.process(AbstractHttp11Processor.java:1082)
	at org.apache.coyote.AbstractProtocol$AbstractConnectionHandler.process(AbstractProtocol.java:623)
	at org.apache.tomcat.util.net.JIoEndpoint$SocketProcessor.run(JIoEndpoint.java:318)
	at java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1145)
	at java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:615)
	at org.apache.tomcat.util.threads.TaskThread$WrappingRunnable.run(TaskThread.java:61)
	at java.lang.Thread.run(Thread.java:745)


```
简单地说　就是不支持的操作符　百思不得其解　最后终于得出两个自己感觉有价值的信息：

```
It means that per the contract of the API, the remove() method is not required to be implemented / available on all implementations. This is because on some implementations, it simply may not make sense - or would be infeasible to implement. A good example would be the iterator returned on a read-only collection - such as the one returned by Collections.unmodifiableCollection().
```

```
Most likely the Collection returned by edgeSet() simply doesn't support the remove() operation, or it does support remove(), yet its Iterator does not.
```
简而言之，就是Java有的集合是没有实现（重写）remove方法，并不适用所有collection，看上面那一段，一个ＱＡ　
Ｑ如下：
```
How to know whether Iterator of a Collection supports remove()?
All Java Collections implement Iterable, so they must provide an Iterator, which specifies an optional method remove(). When remove() is called on the Iterator, it can throw an UnsupportedOperationException.

How do I know whether a Collection in the Java standard library will return an Iterator that supports remove() or not without running code?

Of course I expected this information to be in the Javadoc of the remove() method of the class, but instead found a bunch of links to superclasses and interfaces. For example: http://docs.oracle.com/javase/8/docs/api/java/util/TreeSet.html#iterator-- I did not find any clarification following up on the links, either.
```
其实到此问题已经解决，只需新建一个ＡrrayList来使用Iterator进行遍历，再判断进行remove()即可　如下：

```
for (int i = 0;i<lawyerList.size();i++){
            lawyers.add(lawyerList.get(i));
        }
        Iterator<Lawyer> iterator = lawyers.iterator();
        if("1".equals(customer.getType())){
            while (iterator.hasNext()){
                Lawyer lawyer = iterator.next();
                if(lawyer.getLawyerid().equals(customer.getCustomerid())){
                    iterator.remove();
                }
            }
        }
```



