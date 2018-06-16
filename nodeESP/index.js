const express = require('express');
const bodyParser = require('body-parser');
const path = require('path');
const getRawBody = require('raw-body');
const contentType = require('content-type');

var PORT = process.env.port || 5000;

var app = express()

app.disable('x-powered-by');
app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());
app.use(express.static(path.join(__dirname, 'public')));
app.use(function (req, res, next) 
{
    getRawBody(req, 
    {
        length: req.headers['content-length'],
        limit: '1mb',
        encoding: contentType.parse(req).parameters.charset
    },function (err, string) 
    {
        if (err) 
            return next(err);
        req.text = string;
        next()
    });
});


app.get('/api/name',function(req,res)
{
    res.status(200).send("Esto es un ejemplo");
});

app.post('/api/valor',function(req,res)
{
    res.set('Content-Type', 'text/plain');
    res.status(200).send("<<<"+req.text+">>>");
});

app.listen(PORT,function()
{
    console.log(`Servidor en Marcha http://127.0.0.1:${PORT}`);
});
