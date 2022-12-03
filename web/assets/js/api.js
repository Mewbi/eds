function submit() {
    var name = $("#name").val()
    var last_name = $("#surname").val()
    var full_name = name + " " + last_name
    var msg = $("#message").val()

    if ( msg == "" ){
        alert('Insira uma mensagem')
    }
    if (full_name == " "){
        alert('Insira uma nome')
    }

    $.post("/comment", {
        name: full_name,
        msg: msg
    }, function(response) {
        alert('Obrigado por compartilhar sua opinião conosco.')
    })
}

function populate_comments(response) {
    $("#comments").html('')
    comments = JSON.parse(response)

    for (const i in comments) {
        c = comments[i]

        $("#comments").append(`
            <tr>
                <td>${c.Name}</td>
                <td>${c.Comment}</td>
            </tr>
        `)
    }
}

function get_comments() {
    var auth = $("#auth").val()
    $.post("/comment/view", {
        auth: auth
    }, function(response) {
        if (response != null) {
            populate_comments(response)
        }
    })
}

function get_effectiveness() {
    $.get("/effectiveness", function(response) {
        console.log(response)
    })
}
