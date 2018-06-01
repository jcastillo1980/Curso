const express = require('express');
const sql = require("mssql");
const bodyParser = require('body-parser')
const app = express();

app.use(bodyParser.json())
app.use(bodyParser.urlencoded({extended: true}))

app.get("/mi/",(req,res)=>
{
    res.json(req.body)
});

app.get('/api/:max', function (req, res) 
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

        var server = app.listen(5000, function () {
            console.log('Server is running..');
            console.log("Esto es h y funciona bien, la h funciona bien? h? hhh");
        });

    });
