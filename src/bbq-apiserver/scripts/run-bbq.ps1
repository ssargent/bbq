gci ./secrets | ForEach-Object -Process { $foo = $_.Basename
	$data = gc "./secrets/$_" 
	write-host "$foo - $data"
	[Environment]::SetEnvironmentVariable($foo, $data, "Process")
}

go build

./bbq-apiserver.exe