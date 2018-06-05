const express = require('express');
const sql = require("mssql");
const bodyParser = require('body-parser')
const app = express();
const sqllib = require('./sqllib.js');
const moment = require('moment');


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

app.get('/api/log',function(req,res)
{
    sqllib.insert(sql,"log_procesos",{texto:'algo','st':moment().unix(),'tipo':3},function(err,result)
    {
        res.status(200).send({'err':err,'result':result});
    });
});

    // config for your database
    var config = {
        user: 'externo',
        password: 'externo',
        server: 'oficina.xuitec.com', 
        database: 'lecContador' 
    };

    // connect to your database
    sql.connect(config, function (err) {
    
        if (err) console.log(err);

        var server = app.listen(5000, function () 
        {
            sqllib.setConexion(sql);
            console.log('Server is running..');
            console.log("Esto es h y funciona bien, la h funciona bien? h? hhh");
        });

    });
