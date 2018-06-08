

var sql = null

module.exports.setConexion = function(conexion)
{
    sql = conexion;
}

module.exports.update = function(table,campos,filtro,fun)
{
    let req = new sql.Request();
    let keys = Object.keys(campos);
    txtkeys = "";
    txtfiltros = "";
    primero = true;

    for(i = 0; i < keys.length;i++)
    {
        if(primero == true)
        {
            txtkeys = keys[i]+"=@"+(i+1);
            primero = false; 
        }
        else
        {
            txtkeys = txtkeys + ","+keys[i]+"=@"+(i+1);
        }

        req.input((i+1) + "",campos[keys[i]]);
    }

    if(filtro != null)
    {
        txtfiltros = filtro;
    }

    req.query(texto = "update " + table + " set "+txtkeys + " " + txtfiltros, function (err, resultado) 
    {
        //console.log(texto);
        fun(err, resultado);
    });
}

module.exports.insert = function(tabla, campos, fun)
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
        //console.log(texto);
        fun(err, resultado);
    });

    console.log("texto: ", texto);
}

module.exports.query = function(sqlTexto,campos, fun)
{

    let req = new sql.Request();
    if(campos != null)
    {
        let i = 0;
        for(i = 0; i < campos.length;i++)
        {
            req.input((i+1) + "", campos[i] );
        }
    }

    req.query(sqlTexto, function(err, resultado) 
    {
        console.log(sqlTexto);
        fun(err, resultado);
    });
}