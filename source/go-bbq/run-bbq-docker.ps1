gci ./secrets | ForEach-Object -Process { $foo = $_.Basename
	$data = gc "./secrets/$_" 
	[Environment]::SetEnvironmentVariable($foo, $data, "Process")
}

docker run -t -i -e BBQ_DB_HOST -e BBQ_DB_NAME -e BBQ_DB_PASSWORD -e BBQ_DB_USER -p 21337:21337 myfamilycooks.azurecr.io/bbq/go-bbq:latest