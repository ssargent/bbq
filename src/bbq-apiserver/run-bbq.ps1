gci ./secrets | ForEach-Object -Process { $foo = $_.Basename
	$data = gc "../source/go-bbq/secrets/$_" 
	[Environment]::SetEnvironmentVariable($foo, $data, "Process")
}

go build

./go-bbq.exe