const colors = require('colors');
const path = require('path');
const express = require('express');
const bodyParser = require('body-parser');
const cookieParser = require('cookie-parser');
const methodOverride = require('method-override');
const cluster = require('cluster');
const os = require('os');
var PORT = process.env.port || 5000;

var app = express()


app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());
app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));
app.use(methodOverride('_method'));

app.get('/api/pid',function(req,res)
{
    console.log("cookies:",req.cookies);
    let randomNumber=Math.random().toString();
    res.cookie('cokkieName',randomNumber, { maxAge: 900000, httpOnly: true });
    res.status(200).send({'pid':process.pid,'clusterid':cluster.worker.id});
      
});

if (cluster.isMaster) 
{
    let cpuCount = os.cpus().length;
    for (var i = 0; i < cpuCount; i += 1) 
    {
        cluster.fork();
    }
    cluster.on('online', function(worker) 
    {
        console.log('Worker ' + worker.process.pid + ' is online');
    });

    cluster.on('exit', function(worker, code, signal) 
    {
        console.log('Worker ' + worker.process.pid + ' died with code: ' + code + ', and signal: ' + signal);
        console.log('Starting a new worker');
        cluster.fork();
    });

    setTimeout(function()
    {
        for (const id in cluster.workers) 
        {
            cluster.workers[id].send("tomaaaa");
        }
    },10000);
} 
else 
{
    process.on('message', function(message) 
    {
        console.log("Soy ",process.pid,":",message);
    });

    app.listen(PORT,function()
    {
        console.log(`http://127.0.0.1:${PORT}`);
    });
}