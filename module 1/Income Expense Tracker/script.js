let currentType = 'Income';
const incomeList = [];
const expenseList = [];
var curfilter="showall"

function setDropdownText(type) {
    document.getElementById('type-button').textContent = type;
    currentType = type;
}

function addTask() {
    const description = document.getElementById('description').value.trim();
    const amount = parseFloat(document.getElementById('amount').value.trim());

    if (description === '' || isNaN(amount)) {
        alert('Please enter a valid description and numeric amount.');
        return;
    }

    const newItem = { description, amount, type: currentType };

    if (currentType === 'Income') {
        incomeList.push(newItem);
    } else {
        expenseList.push(newItem);
    }

    showInfo('showall');
    updateTable('showall')
    updateBalance();

    document.getElementById('description').value = '';
    document.getElementById('amount').value = '';
}

function deleteRow(button, description, type) {
    if (type === 'Income') {
        const index = incomeList.findIndex(item => item.description === description);
        if (index !== -1) incomeList.splice(index, 1);
    } else {
        const index = expenseList.findIndex(item => item.description === description);
        if (index !== -1) expenseList.splice(index, 1);
    }
    showInfo('showall')
    updateTable('showall');
    updateBalance();
}

function showInfo(filter) {

    if (curfilter!= filter) {
    var prevwinele = document.getElementById(curfilter);
    prevwinele.classList.remove("bg-success-subtle","border-bottom","border-success"); 
    var curwinele = document.getElementById(filter);
    curwinele.classList.add("bg-success-subtle","border-bottom","border-success"); 
    updateTable(filter);
    curfilter=filter;
    }}

function updateTable(filter) {
    const tableBody = document.getElementById('InfoList');
    tableBody.innerHTML = '';

    let listToShow = [];
    if (filter === 'showincome') {
        listToShow = incomeList;
    } else if (filter === 'showexpense') {
        listToShow = expenseList;
    } else {
        listToShow = [...incomeList, ...expenseList];
    }

    listToShow.forEach(item => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${item.description}</td>
            <td>${item.amount}</td>
            <td>${item.type}</td>
            <td><button class="btn btn-danger btn-sm" onclick="deleteRow(this, '${item.description}', '${item.type}')">X</button></td>
        `;
        tableBody.appendChild(row);
    });
}

function updateBalance() {
    const incomeTotal = incomeList.reduce((sum, item) => sum + item.amount, 0);
    const expenseTotal = expenseList.reduce((sum, item) => sum + item.amount, 0);
    const finalBalance = incomeTotal - expenseTotal;
    document.getElementById('final-balance').textContent = finalBalance;
}