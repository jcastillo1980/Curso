const colors = require('colors');
const http = require('http');
const path = require('path')
const crypto = require('crypto');
const cipher = crypto.createCipher('aes192', 'a password');
const tlog = require('console-timestamp');

const PORT  = process.env.port  || 8080;


console.log(`fichero:[${__filename}]`.blue);
console.log(`directorio:[${__dirname}]`.blue);
console.log(path.join(__dirname,"public"))


var servidor  = http.createServer(function(req,res)
{
    res.end("OK: "+ Date.now()+"\r\n");
    console.log(":" +req.headers['content-length'])

}).listen(PORT,"127.0.0.1",function()
{
    console.log(`Servidor en marcha http://127.0.0.1:${PORT}`.green);
});

servidor.on('error',function()
{
    console.log(`Error abriendo el servidor puerto ${PORT} !!`.red);
    process.exit(-1);
});


let encrypted = '';
cipher.on('readable', () => {
  const data = cipher.read();
  if (data)
    encrypted += data.toString('hex');
});
cipher.on('end', () => {
  console.log(encrypted);
});

cipher.write('some clear text data');
cipher.end();

console.log('hh:mm:ss'.timestamp);

