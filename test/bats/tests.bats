#!/usr/bin/env bats

filler=$PWD/filler

@test "Check that the filler binary is available" {
  run stat $filler
  [ $status -eq 0 ]
}

@test "help" {
  run $filler --help
  [ $status -eq 0 ]
}

@test "version" {
    run $filler version
    [ $status -eq 0 ]
    [[ ${lines[0]} =~ "Version: " ]]
}

@test "filler-run-without-ext" {
    export TEST1="blabla"
    export TEST2="blabla2"
    run $filler --src test/output
    [ $status -eq 0 ]
    run grep "blabla" test/output/a.conf
    [ $status -eq 0 ]
    run grep "blabla2" test/output/b.conf
    [ $status -eq 0 ]
}

@test "filler-run-with-ext" {
    export TEST2="blabla"
    run $filler --src test/output --ext tpl_new
    [ $status -eq 0 ]
    run grep "blabla" test/output/c.conf
    [ $status -eq 0 ]
}

@test "filler-run-with-single-file" {
    export TEST1="blabla"
    run $filler --src test/output/d.conf.tpl_single --ext tpl_single
    [ $status -eq 0 ]
    run grep "blabla" test/output/d.conf
    [ $status -eq 0 ]
}

@test "filler-run-with-missing-var" {
    export TEST1="blabla"
    run $filler --src test/output
    [ $status -ne 0 ]
    [[ ${lines[0]} =~ "ENV variable is missing" ]]
}

@test "filler-run-with-delete" {
    export TEST1="blabla"
    export TEST2="blabla2"
    run $filler --src test/output --delete
    [ $status -eq 0 ]
    run grep "blabla" test/output/a.conf
    [ $status -eq 0 ]
    run grep "blabla2" test/output/b.conf
    [ $status -eq 0 ]
    run ls test/output/*tpl
    [ $status -ne 0 ]
}


