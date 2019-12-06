var pid = 1

function loadFiles() {
    setInterval(() => {
        $.ajax({
            url: "http://localhost:8080/getProcesses",
            success: function(result) {
                let processes = JSON.parse(result)
                let abc = JSON.parse(result)
                $('#divpsbox').empty()

                var i = processes.length
                if (i == 0) {
                    $('#divpsbox').append('<p>collecting data ...</p>');
                }
                $('#divpsbox').append('<table>');
                $('#divpsbox').append('<tr>'); // start table head
                $('#divpsbox').append('<th>' + 'index' + '</th>');
                $('#divpsbox').append('<th>' + 'cpu' + '</th>');
                $('#divpsbox').append('<th>' + 'pid' + '</th>');
                $('#divpsbox').append('<th>' + 'user' + '</th>');
                $('#divpsbox').append('<th>' + 'command' + '</th>');
                $('#divpsbox').append('<th>' + 'kill' + '</th>');
                $('#divpsbox').append('<th>' + 'launch' + '</th>');
                $('#divpsbox').append('</tr>'); // end table head
                processes.forEach(element => {
                    $('#divpsbox').append('<tr>');
                    $('#divpsbox').append('<td>' + i + '</td>');
                    $('#divpsbox').append('<td>' + element['cpu'] + '</td>');
                    $('#divpsbox').append('<td>' + element['pid'] + '</td>');
                    $('#divpsbox').append('<td>' + element['user'] + '</td>');
                    $('#divpsbox').append('<td>' + element['command'] + '</td>');
                    $('#divpsbox').append('<td>' + '<button onclick="killHandler(' + element['pid'] + ')">kill process</button></td>');
                    var cmd = `'${element['command']}'`
                    $('#divpsbox').append('<td>' + '<button onclick="launchHandler(' + cmd + ')">try launch another instance</button></td>');
                    $('#divpsbox').append('</tr>');
                    i--
                });

                $('#divpsbox').append('</table>');
            }
        });
    }, 2000);
}

function killHandler(thisPID) {
    let that = this;
    $.ajax({
        type: 'POST',
        url: "http://localhost:8080/postPid",
        data: { pid: thisPID },
        success: function() {
            that.loadFiles();
        }
    })
};

function launchHandler(thisCOMMAND) {
    let that = this;
    $.ajax({
        type: 'POST',
        url: "http://localhost:8080/postCommand",
        data: { command: thisCOMMAND },
        success: function() {
            that.loadFiles();
        }
    })
}