	/get/schools -> ms de escola
responsible

	/delete/child/school/driver (pra cancelar com o motorista)
	/get/school/driver 
	/get/school/driver -> vai bater no ms de motorista
	/post/school/driver (pra criar uma matricula)

driver
	/get/schools
	/get/sponsors

school
	/get/drivers
	/get/sponsors
	
account-manager
 
	/get/school?driver=x (validar se a escola tem relação com esse driver)
	/get/driver?school=x	
	
banco 
	school-drivers
	record | escola | email | cnpj | driver | email | cnh
	
	account-manager
	record | escola | email | cnpj | driver | email | cnh | responsible | email | cpf | child | rg  
	