let createPassBut=document.getElementById("createPass");
let copyPass=document.getElementById("passwordLabel");




// Generate Password
function generatePassword() {
    var pass= Math.random().toString(36).slice(2) + Math.random().toString(36).slice(2)
    createLabelAndAppendPassword(pass)
}


function copyText() {
    
    var copy=document.getElementsByTagName('pre')[0].innerHTML;
    navigator.clipboard.writeText(copy);
    copyPass.innerHTML = "Copied";
    
    
}

function createLabelAndAppendPassword(generratedPassword) {
    var passwordlabel=document.getElementById("passwordLabel");
    passwordLabel.innerHTML=""
    passwordlabel.innerHTML = generratedPassword;

}

createPassBut.addEventListener("click",generatePassword)
copyPass.addEventListener("click",copyText)