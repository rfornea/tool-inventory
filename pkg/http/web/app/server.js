// server.js
const express = require("express");
let history = require("connect-history-api-fallback");
const serveStatic = require("serve-static");
const path = require("path");
let app = express();
app.use(history());
app.use(serveStatic(path.join(__dirname, "dist")));
const port = process.env.PORT || 80;
app.listen(port);