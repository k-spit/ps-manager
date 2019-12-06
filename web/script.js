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
                $('#divpsbox').append('<th><button onclick="cpuHandler()">' + 'cpu' + '</button></th>');
                $('#divpsbox').append('<th><button onclick="pidHandler()">' + 'pid' + '</button></th>');
                $('#divpsbox').append('<th><button onclick="userHandler()">' + 'user' + '</button></th>');
                $('#divpsbox').append('<th><button onclick="commandHandler()">' + 'command' + '</button></th>');
                $('#divpsbox').append('<th>' + 'kill' + '</th>');
                $('#divpsbox').append('<th>' + 'launch' + '</th>');
                $('#divpsbox').append('</tr>'); // end table head
                processes.forEach(element => {
                    $('#divpsbox').append('<tr>'); // loop start table body
                    $('#divpsbox').append('<td>' + i + '</td>');
                    $('#divpsbox').append('<td>' + element['cpu'] + '</td>');
                    $('#divpsbox').append('<td>' + element['pid'] + '</td>');
                    $('#divpsbox').append('<td>' + element['user'] + '</td>');
                    $('#divpsbox').append('<td>' + element['command'] + '</td>');
                    $('#divpsbox').append('<td>' + '<button onclick="infoHandler(' + element['pid'] + ')">show pstree</button></td>');
                    $('#divpsbox').append('<td>' + '<button onclick="killHandler(' + element['pid'] + ')">kill process</button></td>');
                    var cmd = `'${element['command']}'`
                    $('#divpsbox').append('<td>' + '<button onclick="launchHandler(' + cmd + ')">try launch another instance</button></td>');
                    $('#divpsbox').append('</tr>');
                    i--
                });

                $('#divpsbox').append('</table>'); // loop end table
            }
        });
    }, 2000);
}

function cpuHandler() {
    let that = this;
    $.ajax({
        type: 'POST',
        url: "http://localhost:8080/cpu",
        success: function() {
            // that.loadFiles();
        }
    })
};

function pidHandler() {
    let that = this;
    $.ajax({
        type: 'POST',
        url: "http://localhost:8080/pid",
        success: function() {
            // that.loadFiles();
        }
    })
};

function userHandler() {
    let that = this;
    $.ajax({
        type: 'POST',
        url: "http://localhost:8080/user",
        success: function() {
            // that.loadFiles();
        }
    })
};

function commandHandler() {
    let that = this;
    $.ajax({
        type: 'POST',
        url: "http://localhost:8080/command",
        success: function() {
            // that.loadFiles();
        }
    })
};

function infoHandler(thisPID) {
    let that = this;
    $.ajax({
        type: 'POST',
        url: "http://localhost:8080/infoPid",
        data: { pid: thisPID },
        success: function(result) {
            console.log(result)

            // for loop result
            // append every line to avoid one liner
            $(function() {
                $("#dialog").append('<p>' + result + '</p>');
            });

        }
    })
};

function killHandler(thisPID) {
    let that = this;
    $.ajax({
        type: 'POST',
        url: "http://localhost:8080/postPid",
        data: { pid: thisPID },
        success: function() {
            // that.loadFiles();
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
            // that.loadFiles();
        }
    })
}