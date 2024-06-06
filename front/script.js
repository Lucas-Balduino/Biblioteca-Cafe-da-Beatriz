document.getElementById('wishlistForm').addEventListener('submit', function(event) {
  event.preventDefault();
  
  const nome = document.getElementById('nome').value;
  const autor = document.getElementById('autor').value;
  const qtd = document.getElementById('qtd').value;

  const table = document.getElementById('wishlistTable').getElementsByTagName('tbody')[0];
  const newRow = table.insertRow();

  const cell1 = newRow.insertCell(0);
  const cell2 = newRow.insertCell(1);
  const cell3 = newRow.insertCell(2);
  const cell4 = newRow.insertCell(3);

  cell1.textContent = nome;
  cell2.textContent = autor;
  cell3.textContent = qtd;
  cell4.innerHTML = '<button onclick="removeRow(this)">Remover</button>';

  document.getElementById('wishlistForm').reset();
});

function removeRow(button) {
  const row = button.parentElement.parentElement;
  row.remove();
}
