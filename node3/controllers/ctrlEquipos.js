const express = require('express');



module.exports.routerEquipos = function(cluster)
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

    return router;
};