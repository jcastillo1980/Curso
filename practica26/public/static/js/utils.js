

(function(document,window){

    var mainHead = {tag:"",value:""};
    var wsock = null;
    var wsock_json = false;

    var objeto = {
        setMainHead: function(tag,value)
        {
            mainHead.tag = tag;
            mainHead.value = value;

            return this;
        },
        jsonToQuery: function(obj)
        {
            return Object.keys(obj).map(k => encodeURIComponent(k) + '=' + encodeURIComponent(obj[k])).join('&');
        },
        post: function(xhrIsForm,url,query, contenido,cbOK,cbFAIL)
        {
            if(window.XMLHttpRequest) 
            {
                xhr = new XMLHttpRequest();
            }
            else if(window.ActiveXObject) 
            {
                xhr = new ActiveXObject("Microsoft.XMLHTTP");
            }

            xhr.onreadystatechange = function()
            {
                if(this.readyState == XMLHttpRequest.DONE) 
                {
                    if(this.status == 200)
                    {
                        if(xhrIsForm == true)
                            cbOK(this.responseText);
                        else
                        {
                            objs = JSON.parse(this.responseText);
                            cbOK(objs);
                        }
                    }
                    else
                        cbFAIL();
                }
            };

            if(query != null)
            {
                if(typeof(query) === 'object')
                {
                    url = url + "?" + this.jsonToQuery(query);
                }
                else if(typeof(query) === 'string')
                {
                    url = url + "?" + query;
                }
            }

            xhr.open('POST', url,true);

            if(mainHead.tag.length > 1)
                xhr.setRequestHeader(mainHead.tag,mainHead.value);

            if(xhrIsForm == true)
                xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
            else
                xhr.setRequestHeader("Content-Type", "application/json");


            if(xhrIsForm == true)
                xhr.send(contenido);
            else
                xhr.send(JSON.stringify(contenido));


            return this;
        },
        get: function(xhrIsForm,url,query,cbOK,cbFAIL)
        {
            if(window.XMLHttpRequest) 
            {
                xhr = new XMLHttpRequest();
            }
            else if(window.ActiveXObject) 
            {
                xhr = new ActiveXObject("Microsoft.XMLHTTP");
            }

            xhr.onreadystatechange = function()
            {
                if(this.readyState == XMLHttpRequest.DONE) 
                {
                    if(this.status == 200)
                    {
                        if(xhrIsForm == true)
                            cbOK(this.responseText);
                        else
                            cbOK(JSON.parse(this.responseText));
                    }
                    else
                        cbFAIL();
                }
            };

            if(query != null)
            {
                if(typeof(query) === 'object')
                {
                    url = url + "?" + this.jsonToQuery(query);
                }
                else if(typeof(query) === 'string')
                {
                    url = url + "?" + query;
                }
            }

            xhr.open('GET', url,true);

            if(mainHead.tag.length > 1)
                xhr.setRequestHeader(mainHead.tag,mainHead.value);

            xhr.send(null);

            return this;
        },
        formToJson: function(idform)
        {
            var ob = {};
            var f = document.getElementById(idform);

            if(typeof(f) === 'undefined')
                return ob;

            for(var i = 0; i < 256; i++)
            {
                if(typeof(f[i+'']) === 'undefined')
                    break;
                let elm = f[i+''];
                
                if(elm.type === 'checkbox')
                {
                    ob[elm.name] = elm.checked;
                }
                else if(elm.type === 'radio')
                {
                    if(elm.checked == true)
                        ob[elm.name] = elm.value;
                }
                else
                {
                    ob[elm.name] = elm.value;
                }
            }

            return ob;
        },
        onload: function(cb)
        {
            window.onload = cb;
        },
        open: function(isJSON, uri,cbOPEN,cbCLOSE,cbFAIL,cbMESSAGE) // uri: -> ws://{location.host}}/pepe
        {
            if(wsock == null)
            {
                if(uri.indexOf(":/") < 0)
                {
                    if(location.protocol.indexOf("https:") > 0)
                    {
                        uri = "wss://" + location.host + uri; 
                    }
                    else
                    {
                        uri = "ws://" + location.host + uri; 
                    }
                }
                try
                {
                    wsock = new WebSocket(uri);
                }
                catch(ex)
                {
                    wsock = null;
                    cbFAIL();
                    return;
                }

                if(isJSON == true)
                {
                    wsock_json = true;
                    wsock.onopen = function(){cbOPEN();};
                    wsock.onclose = function(){wsock = null;cbCLOSE();};
                    wsock.onerror = function(){cbFAIL();};
                    wsock.onmessage = function(cont){cbMESSAGE(JSON.parse(cont.data));};
                }
                else
                {
                    wsock_json = false;
                    wsock.onopen = function(){cbOPEN();};
                    wsock.onclose = function(){wsock = null;cbCLOSE();};
                    wsock.onerror = function(){cbFAIL();};
                    wsock.onmessage = function(cont){cbMESSAGE(cont.data);};                   
                }
            }
            else
            {
                cbFAIL();
            }


            return this;
        },
        close: function()
        {
            if(wsock != null)
            {
                wsock.close();
                wsock = null;
            }
            return this;
        },
        send: function(cont)
        {
            if(wsock != null)
            {
                if(wsock_json == true)
                {
                    wsock.send(JSON.stringify(cont));
                }
                else
                {
                    wsock.send(cont);
                }
            }
        }

    }

    if(typeof(window.$$) === 'undefined' )
        window.$$ = objeto;

})(document,window)


