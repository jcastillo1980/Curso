unsigned char antes1  =  Z1;
unsigned char antes2  =  Z2;
unsigned char antes3  =  Z3;

void interrupcion_cn()
{
	if(is1PH == false)
	{
		if(antes1 != Z1)
		{
			antes1 = Z1
			if(escucha1 == true)
			{
			
				if(modo == micro)
				{
					hace cosas, en funcion estado. cuando actuas el igbt lanza el timer y despues no escucha
					el timer vuelve a encender el igbt, pero
				}
				else if(modo == macro)
				{
				}
			}
		}
		
		if(antes2 != Z2)
		{
			antes2 = Z2
			if(escucha2 == true)
			{
			
				if(modo == micro)
				{
				}
				else if(modo == macro)
				{
				}
			}
		}
		
		if(antes3 != Z3)
		{
			antes3 = Z3
			if(escucha3 == true)
			{
			
				if(modo == micro)
				{
				}
				else if(modo == macro)
				{
				}
			}
		}
	}
	else
	{
		if(antes1 != Z1)
		{
			antes1 = Z1
			if(escucha1 == true)
			{
			
				if(modo == micro)
				{
				}
				else if(modo == macro)
				{
				}
			}
		}
	}
	
	flagCN = 0;
}