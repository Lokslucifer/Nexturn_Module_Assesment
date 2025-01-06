var alltask = [];
var task = document.getElementById("task");
var tasklist = document.getElementById("taskList");
var unusedkey = [];
var curwin="showall"
function addtask() {
    console.log("add")
    var newtaskval = task.value;
    if (newtaskval.length > 0) {
        var newtask = document.createElement("li");
        newtask.classList.add("list-group-item", "d-flex", "justify-content-between", "align-items-center", "bg-success-subtle", "rounded", "mb-1");
        newtask.innerHTML = '<div><input type="checkbox" onclick="checkboxclick(this)" class="me-1">' + 
                            '<span id="tasktxt">' + newtaskval + '</span></div>' +
                            '<button class="badge text-bg-danger rounded-pill border border-0" style="width:1.5rem;height:1.5rem;" ' +
                            'onclick="delefunc(this)">X</button>';
        tasklist.appendChild(newtask);
        task.value = "";  // Clear input field after adding task
        if(curwin=="showcom"){
            newtask.classList.add("d-none")
        }
    }
}

function checkboxclick(checkbox) {
    var txtspan = checkbox.parentElement.getElementsByTagName("span")[0];
    if (checkbox.checked) {
        txtspan.style.textDecoration = "line-through";
    } else {
        txtspan.style.textDecoration = "none";
    }
    if(curwin!="showall"){
        if(curwin=='showcur'){
            if(checkbox.checked){
                checkbox.parentElement.parentElement.classList.add('d-none')

            }
        }
        else{
            if(!checkbox.checked){
                checkbox.parentElement.parentElement.classList.add('d-none')

            }

        }
    }
}

function delefunc(button) {
    var par = button.parentElement;
    par.remove();
}

function showinfo(nextwin) {
    console.log(nextwin, curwin);
    console.log("clicked");

    if (curwin != nextwin) {
        var prevwinele = document.getElementById(curwin);
        prevwinele.classList.remove("bg-success-subtle","border-bottom","border-success"); 

        var curwinele = document.getElementById(nextwin);
        curwinele.classList.add("bg-success-subtle","border-bottom","border-success"); 
        var taskelements = tasklist.children; // Ensure tasklist is defined
        console.log("taskelemtns",taskelements.length)
        // Show all tasks if 'all' is selected
        if (nextwin == 'showall') {
            for ( ele of taskelements) {
                ele.classList.remove('d-none')
                 // Use style.display instead of setAttribute
            }
        }
        // Show only uncompleted tasks if 'cur' is selected
        else if (nextwin == "showcur") {

            for (let ele of taskelements) {
                var firstele = ele.firstElementChild;
                if (!firstele.firstElementChild.checked) { // Check if the task is incomplete
                    ele.classList.remove('d-none')
                } else {
                    ele.classList.add('d-none')
                }
            }
        }
        // Show only completed tasks for other cases
        else {

            for (let ele of taskelements) {
                var firstele = ele.firstElementChild;

                console.log("tash")
                if (firstele.firstElementChild.checked) { // Check if the task is 
                 
                    ele.classList.remove('d-none')
                } else {
                    ele.classList.add('d-none')
                }
            }
        }

        curwin = nextwin; // Update curwin to reflect the current window/view
    }
}
