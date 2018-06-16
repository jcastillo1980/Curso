const express = require('express');



module.exports.routerEquipos = function(cluster,sqllib)
{
    var router = express.Router();

    router.get('/pid',function(req,res)
    {
        console.log("cookies:",req.cookies);
        let numeroAccesos = 1;
        if(typeof(req.cookies.numeroAccesos) != 'undefined')
            numeroAccesos = parseInt(req.cookies.numeroAccesos) + 1;
        res.cookie('numeroAccesos',numeroAccesos, { maxAge: 900000, httpOnly: true });
        res.status(200).send({
            'pid':process.pid,
            'clusterid':cluster.worker.id,
            'cookie':req.cookies
        });
    });

    router.get('/kill',function(req,res)
    {
        res.status(200).send({'pid':process.pid,'date':new Date()});
        setTimeout(100,function()
        {
            conosle.log("me han matado a los 100ms");
            process.exit(-1);
        });
    });

    router.get("/mi",(req,res)=>
    {
        res.json(req.body)
    });
    
    
    router.get('/uplog/:id',function(req,res)
    {
        sqllib.update ("log_procesos",{texto:'mas cosas','st':moment().unix(),'tipo':4},`where id='${req.params.id}'`,function(error,respuesta)
        {
            res.status(200).send({'error':error,'resultado':respuesta});
        });
    });
    
    router.get('/log',function(req,res)
    {
        sqllib.insert("log_procesos",{texto:'algo','st':moment().unix(),'tipo':3},function(err,result)
        {
            res.status(200).send({'error':err,'resultado':result});
        });
    });
    
    router.post('/query', function(req,res)
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

    return router;
};