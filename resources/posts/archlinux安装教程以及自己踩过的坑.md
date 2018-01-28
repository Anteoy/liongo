---
date: 2017-09-01 10:19:00
title: archlinux安装教程以及自己踩过的坑
categories:
    - 架构,linux,archlinux
tags:
    - 架构,linux,archlinux
---

### 引言：
　　linux是一种哲学。最近喜欢上了arch的简洁，可高度定制化，滚动更新和设计哲学，准备日常办公从ubuntu转向arch，目前已完成安装，正在使用arch写这篇博客，而事实证明arch确实没让我失望，它的确是一个非常不错的发行版.
### 安装环境：
 1. cpu: i5 ram:12G 台式电脑
 2. 一块硬盘（有剩余空间或新硬盘都可以，我这里用的是1T新硬盘，GPT分区，UEFI启动）
 3. 一块8G U盘
 4. 从官网或者其他镜像源下载的iso系统镜像（e.g:archlinux-2017-08.01-x86_64.iso）
### 安装过程：
#### 制作U盘启动基础系统
1. 在ubuntu或者其他Linux系统中执行
```
fdisk -l
```
查看所有硬盘和分区，记录下U盘的磁盘标识，比如我这里是/dev/sdc
2. 使用dd命令制作U盘安装启动系统
```
dd if=xxx.iso of=/dev/sdc
```
if,of可简记为input file,output file.
注意：如果你的系统和我一样有mbr+bios和gpt+uefi的不同硬盘分区方式及启动方式的话,不推荐使用ultraiso进行刻录，我分别尝试了使用相同镜像，一个U盘使用dd，一个U盘使用ultraiso，ultraiso刻录的U盘并不能正常进入基础安装系统，因为我目前电脑有三块硬盘，其中两块都是使用mbr分区方式使用bios启动，而因为gpt的无主分区数量限制，以及2T硬盘大小限制，以及gpt对磁盘的利用率更高，于是我选择了在此块新硬盘采用gpt的分区方式，使用uefi启动.
#### 硬盘分区
1. 插入U盘开机选择UEFI或传统bios启动
- U盘启动，如果不是UEFI，请选择传统模式的U盘启动，一般是开机按F12,F10,ESC，DELETE等键，我这里是F12.
2. 更新系统时间
    - timedatectl set-ntp true
2. 使用fdisk对硬盘进行分区
##### 附上我这里使用的分区方案（在后面的分区方案中，我取消了usr分区，交给了/）
- /swap 16G
- / 300G
- /home 200G
- /boot/efi 10G
##### 这里使用的分区命令
- fdisk -l #查看当前所有硬盘和分区信息
- fdisk /dev/sda #进入指定硬盘进行操作
- 进入后 w 保存退出 q 不保存退出 n 新建分区 然后选择分区序列号 选择起始扇区，一般前面几个可以直接回车默认，传统硬盘都是 512字节扇区，可根据硬盘说明扇区大小进行计算，或者结尾扇区使用+100G这种形式更加简单。d 删除分区
##### 格式化分区（这里有坑，注意swap分区和/boot/efi分区格式化方式不同）
- fdisk -l
- mkfs.ext4 /dev/sdax （普通分区格式化为ext4）
- mkswap /dev/sdax (swap分区格式化建立方式）
- swapon /dev/sdax（激活系统swap分区）
- mkfs.vfat -F32 /dev/sdaY （boot分区与GPT，UEFI有关，使用此命令格式化为fat32）
##### 分区挂载（这里有坑，注意swap分区是不用挂载的，boot分区挂在应该挂在到/boot/efi,而不是/boot）
- mount /dev/sdax /mnt 根分区
- 使用多个分区，还需要为其他分区创建目录并挂载它们（/mnt/boot、/mnt/home、……）
- mkdir -p /mnt/boot/efi
- mount /dev/sda2 /mnt/boot/efi
- 其他的和上面类似
##### 坑来了，这里一定要看执行下mount看是否成功挂载，否则可能会让你从头再来。
#### 安装
##### 安装基本系统
- pacstrap /mnt base
##### 配置系统（有坑）
- 用以下命令生成 fstab 文件 (用 -U 或 -L 选项设置UUID 或卷标)：
- genfstab -U /mnt >> /mnt/etc/fstab
- 特别提醒：在执行完以上命令后，用cat检查一下生成的 /mnt/etc/fstab 文件是否正确。对比blkid命令下硬盘分区UID和此文件是否对应，我这里就是因为没有mount好，生成的fstab也不对，导致安装完成无法启动.
##### Change root 到新安装的系统：
- arch-chroot /mnt
##### 设置时区
- ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
- 设置时间标准 为 UTC，并调整 时间漂移:
- hwclock --systohc --utc
##### Locale本地化配置
- pacman -S vim (习惯vim直接装一个，使用默认vi也可以)
- vim /etc/locale.gen 打开注释
```
en_US.UTF-8 UTF-8
zh_CN.UTF-8 UTF-8
zh_TW.UTF-8 UTF-8
```
- 接着执行locale-gen以生成locale讯息：
- locale-gen
- 将系统 locale 设置为en_US.UTF-8：
-  echo LANG=en_US.UTF-8 > /etc/locale.conf
##### 设置主机名
- echo archlinux > /etc/hostname
##### 创建一个初始 RAM disk：
- mkinitcpio -p linux（有坑，官方wifi说改动了mkinitcpio.conf可以不用执行这个init，我开始以为我没有修改，不用init，结果一下就踩进去了，友情提示这个最好不要忘了执行init一次）
##### 设置 root 密码:
- passwd
##### 配置网络
 - pacman -S dialog wpa_supplicant netctl wireless_tools #现在不安装 重启之后如果只有wifi则可能无法连接网络
 - 查看网卡名:
 - ip link show
 - 设置启动dhcp:
 - systemctl enable dhcpcd@enp0s2.service
#### 安装引导程序
#####我这里选择的grub，注意这里一定得装，否则是无法引导系统的，并且要特别小心，否则很容易无法对系统进行引导启动.
- UEFI版本：
- pacman -S grub-efi-x86_64
- EFI管理器：
- pacman -S efibootmgr
- 双系统必需管理器：（我这里由于bios和uefi方式不同，无法和ubuntu相互引导）
- 安装进EFI分区：
- grub-install --efi-directory=/boot/efi --bootloader-id=grub
- os-prober 识别硬盘上其他系统的工具:(uefi,bios冲突不能互相引导)：
- pacman -S grub
- grub-install --recheck /dev/sda
- 生成配置文件：
- grub-mkconfig -o /boot/grub/grub.cfg
- 友情提示：生成完成cat下/boot/grub/grub.cfg文件是否正常生成.如果不对需要进行自行检测，正常才能继续下面的操作.
#### 退出chroot模式，并umount
- 先umount /mnt里面的boot,home等分区，然后umount /mnt 根分区
##### 坑来了，到这里基本已完成基础系统的安装，但是注意最好不要使用root然后在关机的瞬间把u盘拔掉，我开始就是，一切正常，但是使用reboot并立马拔掉U盘，导致无法进入系统，这里可能是shutdown的时候有部分文件未写入完成就拔掉U盘可能导致数据异常，所以不能进入系统.于是后来就学乖了，先shutdown now关机，再开机进入，一切正常。
- 友情提示：如果你和我一样同时存在mbr,bios和gpt,uefi，需要进bios设置使用bios还是uefi来进行引导启动，否则无法进入系统.
##### 进入安装好的基础系统，然后依次检查网络连接，ip addr,ping,curl,检查分区及目录fdisk -l,du -h ,df -h，新建sudo用户，这里说几个较重要的
#### 安装字体
- pacman -S wqy-zenhei wqy-microhei （中文字体）
- pacman -S ttf-dejavu pacman -S adobe-source-code-pro-fonts （等宽字体）
- pacman -S wqy-microhei wqy-zenhei
- pacman -S fcitx fcitx-im fcitx-googlepinyin 输入法
- /etc/profile加入：(我这里.xinitrc .xprofile  不会生效，如果你也和我一样不生效，可以参考下我这里)
```
export XIM=fcitx
export XIM_PROGRAM=fcitx
export GTK_IM_MODULE=fcitx
export QT_IM_MODULE=fcitx
export XMODIFIERS=@im=fcitx
```
#### 安装网络管理器
- pacman -S networkmanager
- pacman -S network-manager-applet xfce4-notifyd
- pacman -S network-manager-applet xfce4-notifyd
- networkmanager-pptp
- pacman -S networkmanager-pptp
- systemctl start NetworkManager
- systemctl enable NetworkManager
- 友情提示：注意大小写
#### 安装gnome(根据需要你也可以选择kde,xfce等等其他发行版，我习惯用gnome)
- Intel集成显卡驱动：
- pacman -S xf86-video-intel
- 安装显卡驱动：
- pacman -S xf86-video-vesa
- xorg服务：
- pacman -S xorg-xinit xorg-server
- pacman -S xorg-server xorg-xinit xorg-utils xorg-server-utils
- gnome：
- pacman -S gnome
- lib256选择1 根据自己选择配置 我这里选择的1
- gnome 的窗口管理器：
- pacman -S gdm
- systemctl enable gdm
##### 然后reboot重启就可以进入系统了,安装完成.安装过程参考了下面两位大大@禾白小三飘@u012619242的宝贵经验，结合官方wifi进行安装，在此表示感谢，同时也希望我的这篇博文能帮助大家，少踩坑，如果有问题也可以在下面给我留言，欢迎讨论。
### 参考文献：
1. [丰富的arch wifi资料库](https://wiki.archlinux.org/index.php/)
2. [http://www.linuxidc.com/Linux/2016-09/134953.htm](http://www.linuxidc.com/Linux/2016-09/134953.htm)
3. [VirtualBOX安装Archlinux过程](http://www.jianshu.com/p/b66d14dcaffe)
