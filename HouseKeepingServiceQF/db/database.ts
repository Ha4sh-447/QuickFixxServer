// import mysql from "mysql";

import { createConnection } from "mysql2";

// const connect = mysql.createConnection({
// 	host: "localhost",
// 	port: 3306,
// 	database: "servicesdb",
// 	user: "root",
// 	password: "root@123456",
// });

const connect = createConnection({
	host: "localhost",
	port: 3306,
	database: "servicesdb",
	user: "root",
	password: "root@123456",
});

export default connect;
