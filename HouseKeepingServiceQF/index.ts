import express from "express";
import connect from "./db/database";

const app = express();
const port = process.env.PORT || "9776";

app.get("/", async (req, res) => {
	res.status(200).json({
		message: "TS server 123",
	});
});

app.get("/data", async (req, res) => {
	// const query = "SELECT * FROM `t_usersdb`";
	const query: string = "SELECT * FROM `t_usersdb`";
	connect.query(query, (err, result, fields) => {
		if (err instanceof Error) {
			console.log(err);
			res.status(500).json({
				message: err,
			});
		}

		console.log(result);
		console.log(fields);
		res.status(200).json({
			result: result,
			field: fields,
		});
	});
});

app.listen(port, () => {
	console.log(`Server running on port: ${port}`);
});
