

module.exports.insert = function(sql, tabla, campos, fun)
{
    let req = new sql.Request();
    let keys = Object.keys(campos);
    txtkeys = "";
    txtvalores = "";
    primero = true;

    for(i = 0; i < keys.length; i++)
    {
        if(primero == true)
        {
            txtkeys = keys[i];
            txtvalores = "@"+(i+1);
            primero = false;
        }
        else
        {
            txtkeys = txtkeys + ","+keys[i];
            txtvalores = txtvalores + ",@"+(i+1);
        }

        req.input((i+1) + "",campos[keys[i]]);
    }
    

    req.query(texto = "insert into " + tabla + " ("+txtkeys+") values ("+txtvalores+"); SELECT SCOPE_IDENTITY() AS id;", function (err, resultado) 
    {
        fun(err, resultado);
    });

    console.log("texto: ", texto);
}