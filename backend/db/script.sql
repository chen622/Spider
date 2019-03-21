create table if not exists bilibili_up
(
  mid       int               not null
    primary key,
  name      varchar(255)      null,
  last_time int(20) default 0 not null
);

create table if not exists bilibili_video
(
  mid         int          null,
  aid         int          not null
    primary key,
  title       varchar(255) null,
  description varchar(255) null,
  pic         varchar(255) null,
  created     int(20)      null,
  constraint bilibili_video_bilibili_up_mid_fk
    foreign key (mid) references bilibili_up (mid)
      on delete cascade
);

create table if not exists user
(
  id       int auto_increment
    primary key,
  nickname varchar(64)  not null,
  password varchar(255) not null,
  mail     varchar(255) not null
);

create table if not exists user_bilibili_up
(
  user_id        int not null,
  bilibili_up_id int not null,
  primary key (bilibili_up_id, user_id),
  constraint user_bilibili_up_bilibili_up_mid_fk
    foreign key (user_id) references bilibili_up (mid)
      on delete cascade,
  constraint user_bilibili_up_user_id_fk
    foreign key (user_id) references user (id)
      on delete cascade
);


