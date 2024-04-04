"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const database_1 = __importDefault(require("./db/database"));
const app = (0, express_1.default)();
const port = process.env.PORT || "9776";
app.get("/", async (req, res) => {
    res.status(200).json({
        message: "TS server 123",
    });
});
app.get("/data", async (req, res) => {
    // const query = "SELECT * FROM `t_usersdb`";
    const query = "SELECT * FROM `t_usersdb`";
    database_1.default.query(query, (err, result, fields) => {
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
