---
date: 2016-07-19 22:03:00
title: spring AOP 动态代理 jkd动态代理和cglib动态代理 hibernate使用cglib延迟加载
categories:
    - java
tags:
    - spring
---

spring 的AOP 实现 可以使用jdk的动态代理，也可以使用cglib的动态代理 先说下两者区别：
静态代理：代理之前就已经知道了代理者和被代理者
动态代理：代理之前并不清楚，在运行时使用反射机制动态生成代理类的字节码 无需我们手动编写它的源代码
	jdk动态代理：java.lang.reflect 包中的Proxy类，InvocationHandler 接口提供了生成动态代理类的能力。它必须有被代理对象的接口和实现类，ciglib不需要接口，简单的说 jkd动态代理针对接口，而cglib动态代理针对的是类 比JDK快，但加载cglib的时间比jdk反射的时间长，开发的过程中，如果是反复动态生成新的代理类推荐用 JDK 自身的反射，反之用 cglib。

spring AOP的实现 既可以编写接口使用jdk动态代理实现，也可以使用cglib来进行实现

对于hibernate使用cglib延迟加载，在load（)的时候，cglib会生成一个数据库实体类的代理类，此代理类开始的时候里面的属性值是null，只有当真正请求数据,如user.getName(）的时候，这个代理类就会去数据库里查询出具体的值(通过cglib的回调机制)，然后加载到代理类对应的属性中，最后返回给给调用者。

以下为一个转载的实例：
[](http://www.blogjava.net/Good-Game/archive/2007/11/05/158192.html)

```
/**
*使用类
*/
public class MyClass {
public void method() {
        System.out.println("MyClass.method()");
    }
public void method2() {
        System.out.println("MyClass.method2()");
    }
}
/**
*使用代理
*/
import java.lang.reflect.Method;
import net.sf.cglib.proxy.Enhancer;
import net.sf.cglib.proxy.MethodProxy;
import net.sf.cglib.proxy.MethodInterceptor;
public class Main {
public static void main(String[] args) {
        Enhancer enhancer = new Enhancer();
//在这代理了
        enhancer.setSuperclass(MyClass.class);
        enhancer.setCallback( new MethodInterceptorImpl() );
// 创造 代理 （动态扩展了MyClass类）
        MyClass my = (MyClass)enhancer.create();
        my.method();
    }
private static class MethodInterceptorImpl implements MethodInterceptor {
public Object intercept(Object obj,
                                Method method,
                                Object[] args,
                                MethodProxy proxy) throws Throwable {
            System.out.println(method);
            proxy.invokeSuper(obj, args);
return null;
        }
    }
}
/**
*添加方法过滤
*/
import java.lang.reflect.Method;
import net.sf.cglib.proxy.Enhancer;
import net.sf.cglib.proxy.MethodProxy;
import net.sf.cglib.proxy.MethodInterceptor;
import net.sf.cglib.proxy.NoOp;
import net.sf.cglib.proxy.Callback;
import net.sf.cglib.proxy.CallbackFilter;
public class Main2 {
public static void main(String[] args) {
        Callback[] callbacks =
new Callback[] { new MethodInterceptorImpl(),  NoOp.INSTANCE };
        Enhancer enhancer = new Enhancer();
        enhancer.setSuperclass(MyClass.class);
        enhancer.setCallbacks( callbacks );
//添加 方法过滤器  返回1为不运行 2 为运行
        enhancer.setCallbackFilter( new CallbackFilterImpl() );
        MyClass my = (MyClass)enhancer.create();
        my.method();
        my.method2();
    }
private static class CallbackFilterImpl implements CallbackFilter {
public int accept(Method method) {
if ( method.getName().equals("method2") ) {
return 1;
            } else {
return 0;
            }
        }
    }
private static class MethodInterceptorImpl implements MethodInterceptor {
public Object intercept(Object obj,
                                Method method,
                                Object[] args,
                                MethodProxy proxy) throws Throwable {
            System.out.println(method);
return proxy.invokeSuper(obj, args);
        }
    }
}
```
下面贴下在imooc上自己写的几个小例子：
具体可点击：
[](http://www.imooc.com/learn/214)


jdk动态代理的简单实现：

```
package proxy.dynamicProxy.jdk;

import java.util.Random;

/**
 * 需要代理的具体类继承接口
 */
public class Car implements Moveable {

	@Override
	public void move() {
		//实现开车
		try {
			Thread.sleep(new Random().nextInt(1000));
			System.out.println("汽车行驶中....");
		} catch (InterruptedException e) {
			e.printStackTrace();
		}
	}

}
########################################
package proxy.dynamicProxy.jdk;

/**
 *jdk实现动态代理模式，必须使用接口
 */
public interface Moveable {
	void move();
}
########################################
package proxy.dynamicProxy.jdk;


import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;

/**
 * 代理者，必须集成InvocationHandler 实现invoke方法
 */
public class TimeHandler implements InvocationHandler {

	public TimeHandler(Object target) {
		super();
		this.target = target;
	}

	//代理对象
	private Object target;

	/*
	 * 只能代理实现了接口的类，没有实现接口的类不能实现JDK动态代理
	 * 参数：
	 * proxy  被代理对象
	 * method  被代理对象的方法
	 * args 方法的参数
	 *
	 * 返回值：
	 * Object  方法的返回值
	 * */
	@Override
	public Object invoke(Object proxy, Method method, Object[] args)
			throws Throwable {
		long starttime = System.currentTimeMillis();
		System.out.println("汽车开始行驶....");
		method.invoke(target);
		long endtime = System.currentTimeMillis();
		System.out.println("汽车结束行驶....  汽车行驶时间："
				+ (endtime - starttime) + "毫秒！");
		return null;
	}

}
#####################################
package proxy.dynamicProxy.jdk;

import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Proxy;


public class Test {

	/**
	 * JDK动态代理测试类
	 */
	public static void main(String[] args) {
		Car car = new Car();
		InvocationHandler h = new TimeHandler(car);
		Class<?> cls = car.getClass();
		/**
		 * loader  类加载器
		 * interfaces  实现接口
		 * h InvocationHandler 调用处理者
		 */
		Moveable m = (Moveable)Proxy.newProxyInstance(cls.getClassLoader(),
												cls.getInterfaces(), h);
		m.move();
	}

}

```
cglib动态代理简单实现：

```
package proxy.dynamicProxy.cglib;

public class Train {

	public void move(){
		System.out.println("火车行驶中...");
	}
}
##########################################
package proxy.dynamicProxy.cglib;

import net.sf.cglib.proxy.Enhancer;
import net.sf.cglib.proxy.MethodInterceptor;
import net.sf.cglib.proxy.MethodProxy;

import java.lang.reflect.Method;

public class CglibProxy implements MethodInterceptor {

	private Enhancer enhancer = new Enhancer();//实例化强化剂对象

	public Object getProxy(Class clazz){
		//设置创建子类的类
		enhancer.setSuperclass(clazz);
		enhancer.setCallback(this);

		return enhancer.create();
	}

	/**
	 * 针对类来实现代理，先为目标类生产一个子类，通过方法拦截技术拦截所有父类方法的调用
	 * 拦截所有目标类（父类）方法的调用
	 * obj  目标类的实例
	 * m   目标方法的反射对象
	 * args  方法的参数
	 * proxy代理类的实例
	 */
	@Override
	public Object intercept(Object obj, Method m, Object[] args,
			MethodProxy proxy) throws Throwable {
		System.out.println("日志开始...");
		//代理类调用父类的方法
		proxy.invokeSuper(obj, args);
		System.out.println("日志结束...");
		return null;
	}

}
##################################
package proxy.dynamicProxy.cglib;

import proxy.dynamicProxy.cglib.CglibProxy;

public class Client {

	/**
	 * @param args
	 */
	public static void main(String[] args) {

		CglibProxy proxy = new CglibProxy();
		Train t = (Train)proxy.getProxy(Train.class);
		t.move();
	}

}


```


