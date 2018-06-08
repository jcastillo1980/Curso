const express = require('express');
const sql = require("mssql");
const bodyParser = require('body-parser')
const app = express();
const sqllib = require('./sqllib.js');
const moment = require('moment');

const PORT = process.env.port || 5000;


app.use(bodyParser.json())
app.use(bodyParser.urlencoded({extended: true}))

app.get("/mi/",(req,res)=>
{
    res.json(req.body)
});

app.get('/api/list/:max', function (req, res) 
{
   
    let re = new sql.Request();
    re.input("1",req.params.max);
    // query to the database and get the records
    re.query('select * from log_procesos where id > @1 order by id asc', function (err, recordset) {
        
        if (err) console.log(err)

        // send records as a response
        res.status(200).send({'len':recordset.recordset.length,'dato0':recordset.recordset[0]})
        
    });

    
});

app.get('/api/uplog/:id',function(req,res)
{
    sqllib.update ("log_procesos",{texto:'mas cosas','st':moment().unix(),'tipo':4},`where id='${req.params.id}'`,function(error,respuesta)
    {
        res.status(200).send({'error':error,'resultado':respuesta});
    });
});

app.get('/api/log',function(req,res)
{
    sqllib.insert("log_procesos",{texto:'algo','st':moment().unix(),'tipo':3},function(err,result)
    {
        res.status(200).send({'error':err,'resultado':result});
    });
});

app.post('/api/query', function(req,res)
{
    if(typeof(req.body.sqltexto) == 'undefined')
    {
        res.status(200).send({'error':'no existe campo post sqltexto'});
    }
    else
    {
        sqllib.query(req.body.sqltexto,null,function(err,result)
        {
            res.status(200).send({'error':err,'resultado':result});
        });
    }
});


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

    var server = app.listen(PORT, function () 
    {
        sqllib.setConexion(sql);
        console.log(`Servidor en Marcha http://127.0.0.1:${PORT}`);
    });

});
