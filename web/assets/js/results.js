$(document).ready(function () {

    populate_effectiveness()
});

function populate_effectiveness() {
    $.get("/effectiveness", function(response) {

    $("#questions").html('')
    effectiveness = JSON.parse(response)
    
    for (const q in effectiveness) {
        e = effectiveness[q]
        $("#questions").append(`
            <tr>
                <td>${q}</td>
                <td>
                    <p><b>Vezes escolhidas:</b> ${e.Total}</p>
                    <br>
                    <p><b>Resultado correto com diagnóstico clínico:</b> ${e.Correct}</p>
                </td>
                <td>
                    <canvas id="chart-${q}" style="width:100%;max-width:300px"></canvas>
                </td>
            </tr>
        `)

        var correct = Number((e.Effectiveness).toFixed(4)) * 100
        var incorret = 100 - correct
        var xValues = ["Corretos", "Incorretos"];
        var yValues = [correct, incorret];
        var barColors = [
            "#33fff6",
            "#ff5757",
        ];

        new Chart("chart-"+q, {
            type: "doughnut",
            data: {
                labels: xValues,
                datasets: [{
                    backgroundColor: barColors,
                    data: yValues
                }]
            },
            options: {
                title: {
                    display: true,
                    text: "Taxa de assertividade da questão "+ q
                }
            }
        });
    }

    })
}

function populate_data() {
    var auth = $("#auth").val()
    $.post("/effectiveness/populate", {
        auth: auth
    }, function(result) {
        $("#questions").html('')
        populate_effectiveness()
    })
}

