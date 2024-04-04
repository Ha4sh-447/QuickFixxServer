"use strict";
// import mysql from "mysql";
Object.defineProperty(exports, "__esModule", { value: true });
const mysql2_1 = require("mysql2");
// const connect = mysql.createConnection({
// 	host: "localhost",
// 	port: 3306,
// 	database: "servicesdb",
// 	user: "root",
// 	password: "root@123456",
// });
const connect = (0, mysql2_1.createConnection)({
    host: "localhost",
    port: 3306,
    database: "servicesdb",
    user: "root",
    password: "root@123456",
});
exports.default = connect;
