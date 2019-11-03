package main

/*
// 建表

CREATE TABLE `userinfo` (
    `uid` INTEGER PRIMARY KEY AUTOINCREMENT,
    `username` VARCHAR(64) NULL,
    `departname` VARCHAR(64) NULL,
    `created` DATE NULL
);

CREATE TABLE `userdeatail` (
    `uid` INT(10) NULL,
    `intro` TEXT NULL,
    `profile` TEXT NULL,
    PRIMARY KEY (`uid`)
);

 */

// 例子中使用了支持 database/sql 接口的 https://github.com/mattn/go-sqlite3
// 因此代码和使用 MySql 的几乎一样，只是 driver 不同
