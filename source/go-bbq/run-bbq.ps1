gci ./secrets | ForEach-Object -Process { $foo = $_.Basename
    $data = gc "./secrets/$_" 
    [Environment]::SetEnvironmentVariable($foo, $data, "Process")
}

go build

./go-bbq.exe