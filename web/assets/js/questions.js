function get_questions() {
    $.get("/questions", function(response) {
        $("#questions-test").html('')
        var questions = JSON.parse(response)
    
        for (const i in questions) {
            var n = i + 1
            q = questions[i]
            console.log(q)
            $("#questions-test").append(`
                <tr>
                    <td>${n}º</td>
                    <td>
                        ${q.content}
                    </td>
                </tr>
            `)
        }

        $("#questions-test").append(`
            <tr>
                <td colspan="2">
                    <button type="submit" class="mt-3 btn btn-primary">Enviar</button>
                </tr>
            </tr>
        `)
    })
}

