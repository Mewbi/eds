#!/usr/bin/env bash

if [[ -z "$(type -P sqlite3)" ]]; then
    echo "Install sqlite3 to use this script"
    exit 1
fi

function help() {
    cat << END
    Usage: ./db.sh <option>

    Options:

    --create|-c)
        Create new database based in schema file

    --populate|-p)
        Populate database based in data file
END
}

function create() {
    sqlite3 database.db < migrations/sqlite/schema.sql
}

function populate() {
    n=100
    file="migrations/sqlite/data.sql"
    questions=("a1" "a2" "a3" "a4" "a5" "a6" "a7" "a8" "a9" "a10" "b1" "b2" "b3" "b4" "b5" "b6" "b7" "b8" "b9" "b10" "c1" "c2" "c3" "c4" "c5" "c6" "c7" "c8" "c9" "c10")
    names=("James" "Stewart" "John" "Jessie" "Christian" "Josef" "Steven" "Mark" "Huston" "Halliday")
    if [[ -n $1 ]]; then
        n=$1
    fi

    echo "INSERT INTO test_results (id, name, email, responses, created_at, confirmation) VALUES " > $file
    for i in $(seq $n); do
        id="'$(echo $RANDOM | md5sum | head -c 10)'"
        name="${names[ $(expr $RANDOM % ${#names[@]}) ]}"
        email="'${name,,}@gmail.com'"
        name="'${name}'"

        responses="'["
        for j in $(seq 10); do
            question="${questions[ $(expr $RANDOM % ${#questions[@]}) ]}"
            status="true"
            if (( $RANDOM % 2 )); then
                status="false"
            fi
            response="{\"question_id\": \"${question}\", \"response\": ${status}}"
            if (( j < 10 )); then
                response="${response}, "
            fi
            responses="${responses}${response}"
        done
        responses="${responses}]'"

        date="'DATE()'"
        confirmation="true"
        if (( $RANDOM % 2 )); then
            confirmation="false"
        fi

        data="(${id}, ${name}, ${email}, ${responses}, ${date}, ${confirmation})"

        if (( i < $n )); then
            data="${data}, "
        else
            data="${data};"
        fi
        echo $data >> $file
    done
    sqlite3 database.db < $file
}

case ${1,,} in 
    -c|--create)
        create
    ;;
    -p|--populate)
        shift
        populate $1
    ;;
    *)
        help
    ;;
esac
