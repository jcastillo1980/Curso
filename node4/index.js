const colors = require('colors');
const cluster = require('cluster');
const os = require('os');


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

    /*setTimeout(function()
    {
        for (const id in cluster.workers) 
        {
            cluster.workers[id].send("tomaaaa");
        }
    },10000);*/
} 
else 
{
    process.on('message', function(message) 
    {
        console.log("Soy ",process.pid,":",message);
    });

    const path = require('path');
    const express = require('express');
    const bodyParser = require('body-parser');
    const cookieParser = require('cookie-parser');
    const methodOverride = require('method-override');
    const sqllib = require('./lib/sqllib.js');
    const sql = require("mssql");
    
    var PORT = process.env.port || 5000;

    var app = express()
    
    app.disable('x-powered-by');
    app.use(bodyParser.urlencoded({ extended: false }));
    app.use(bodyParser.json());
    app.use(cookieParser());
    app.use(express.static(path.join(__dirname, 'public')));
    app.use(methodOverride('_method'));

    // config for your database
    var config = {
        user: 'externo',
        password: 'externo',
        server: 'oficina.xuitec.com', 
        database: 'lecContador' 
    };

    // connect to your database
    sql.connect(config, function (err) 
    {

        if (err)
        {
            console.log(err);
            process.exit(-1);
        } 

        app.listen(PORT, function () 
        {
            sqllib.setConexion(sql);
            app.use('/api',require('./controllers/ctrlEquipos.js').routerEquipos(cluster,sqllib));
            console.log(`Servidor en Marcha http://127.0.0.1:${PORT}`);
        });

    });

}