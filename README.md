# CCM(Not Finished)

> CCM is CCM Cloud Messaging. It use to push subscription message to android phone.
>


[Official Website: ccm.ink](https://ccm.ink)

Because the Google.Inc pulled Out Of China, Android phone of China don't have a unified message pushing service like [GCM](https://developers.google.com/cloud-messaging/), [FCM](https://firebase.google.com/docs/cloud-messaging/). So lots of China Android Apps try to use some not elegant ways. Just like the apps from Alimama will awaken each other to keep their run at the backend and push messages to the user.

![1553158548836.png](https://i.loli.net/2019/05/26/5cea7459b3f3592516.png)
<p style="text-align: center;font-style: italic;">How to use GCM to push message</p>

More and more App try to use this way to keep live and Android phone has been lagging and need bigger battery. So the Chinese mobile phone company use lots of ways to kill these App. At the end of the war, the phone company won but the Android users also lost the timely notification service.

These project is try to provide a way to help Android user to subscribe some Youtuber or Weibo from someone they like.

These Project use email to push message, and the backend is use Golang and Iris Framework to build.

------

> CCM 是 CCM 消息推送服务的缩写，它用来帮助安卓用户去订阅喜欢的B站Up主或微博博主。

因为谷歌退出中国市场的原因，安卓手机在中国失去了统一的推送服务，例如：GCM。因此许多中国开发者尝试用一些笨拙的方法来实现消息的推送，比如阿里全家桶、某度全家桶等会使用链式启动的方式进行互相唤醒。

![1553158548836.png](https://i.loli.net/2019/05/26/5cea7459b3f3592516.png)
<p style="text-align: center;font-style: italic;">如何使用 GCM 来推送消息</p>

越来越多这样的应用的出现，导致安卓手机变得越来越卡顿并且费电。于是中国的手机厂商开始也使用各式各样的方法来对这些应用进行限制。最后，往往都是手机厂商获得了胜利，同时安卓用户也失去了及时的推送服务。

这个项目便是帮助安卓用户去订阅自己喜欢的Up主或微博博主，项目是使用邮件的方式来推送消息的，同时该项目是使用 Golang 与 Iris 框架编写的。