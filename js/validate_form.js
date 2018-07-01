(function() {
  'use strict';
  window.addEventListener('load', function() {
    var form = document.getElementById('newBonsai');
    form.addEventListener('submit', function(event) {

      if (validateBonsaiForm() == false) {
        event.preventDefault();
        event.stopPropagation();
      }

      form.classList.add('was-validated');
    }, false);
  }, false);
})();

function validateBonsaiForm() {

 var result = true;

 if (document.getElementById('imgInp').value == "") {
   result = false;
 }

 if (!isPositiveNumber(document.getElementById('age').value)) {
   alert("Invalid number for bonsai age");
   result = false;
 }

 if (document.getElementById('name').value == "") {
   result = false;
 }

 if (document.getElementById('type').value == "Choose type") {
   result = false;
 }

 if (document.getElementById('species').value == "Choose species") {
   result = false;
 }

 if (document.getElementById('style').value == "Choose style") {
   result = false;
 }

 return result;

}

function isPositiveNumber(n) {
  return !isNaN(parseFloat(n)) && isFinite(n) && n>0;
}
