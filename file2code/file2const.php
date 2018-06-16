<?php

    echo "Programa de conversión\r\n";
    if(count($argv) < 2)
    {
        echo "Error en argumentos !! \r\n";
        exit(-1);
    }

    function contentFile($name)
    {
        $fd = fopen($name, "rb");
        if($fd === false)
        {
            return null;
        }
        $contenido = fread($fd, filesize($name));
        fclose($fd);

        $retorno = "";
        $hay_espacio = false;
        for($I=0;$I<strlen($contenido);$I++)
        {
            $valor = ord($contenido[$I]);
            if(($valor != 10) && ($valor != 13))
            {
                if($contenido[$I] == ' ')
                {
                    if($hay_espacio == true)
                    {
                        $hay_espacio = true;
                    }
                    else
                    {
                        $hay_espacio = true;
                        $retorno = $retorno.$contenido[$I];
                    }
                }
                else
                {
                    $hay_espacio = false;
                    $retorno = $retorno.$contenido[$I];
                }
            }
        }

        return $retorno;
    }

    $fd = fopen('out.h', 'w');

    if($fd === false)
    {
        echo "No se puede crear out.h !!\r\n";
        exit(-1);
    }

    for($I = 1; $I < count($argv); $I++)
    {
        $contenido = contentFile($argv[$I]);
        if($contenido == null)
        {
            echo "{$argv[$I]}:ERROR\r\n";
        }
        else
        {
            $nameVar = str_replace(".","_",$argv[$I]);
            fprintf($fd,"const char {$nameVar}[] = {");
            $primero = true;
            for($Y = 0; $Y < strlen($contenido); $Y++)
            {
                if($primero == true)
                {
                    $primero = false;
                    fprintf($fd,"0x%02X",ord($contenido[$Y]));
                }
                else
                {
                    fprintf($fd,",0x%02X",ord($contenido[$Y]));
                }
            }
            fprintf($fd,"};\r\n");
            echo "{$argv[$I]}:OK\r\n";
        }
    }

    fprintf($fd,"\r\n");

    fclose($fd);

?>