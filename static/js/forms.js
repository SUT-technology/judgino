let allCheckInputs = document.getElementsByClassName("check-input");

const getValidations = (inp) => {
  let validationsSplit = inp.dataset.validations.split(",");
  let validations = {
    max: inp.dataset.max,
    min: inp.dataset.min,
  };
  validationsSplit.forEach((v) => {
    validations[v] = true;
  });
  return validations;
};
for (let i = 0; i < allCheckInputs.length; i++) {
  let inpBox = allCheckInputs[i];
  let inpCon = inpBox.getElementsByClassName("inp-con")[0];
  inpCon.setAttribute("class","inp-con w-full h-12 rounded-lg border border-2 border-purple-700 bg-opacity-30 backdrop-blur-4xl p-2 flex items-center justify-between")
  let inp = inpBox.getElementsByClassName("text-input")[0];
  inp.setAttribute(
    "class",
    "w-full border-none bg-transparent outline-none text-sm pt-2 ml-auto text-input"
  );
  let validations = getValidations(inp);
  let errorMsg = "";
  let label = inpBox.getElementsByTagName("label")[0];
  let validationsSplit = inp.dataset.validations.split(",");
  if (validationsSplit.includes("required")) {
    label.textContent += " *";
  }
  inp.onchange = () => {
    if (inp.value.length > 0) label.classList.add("active");
    else label.classList.remove("active");
  };
  inp.onkeyup = () => {
    errorMsg = checkInput(inp.value, validations, label.textContent);
    inpBox.getElementsByClassName("error-message")[0].textContent = errorMsg;
    let errorIcon = inpBox.getElementsByClassName("error-icon")[0];
    let successIcon = inpBox.getElementsByClassName("success-icon")[0];
    if (errorMsg) {
      inpCon.classList.remove("border-purple-700");
      inpCon.classList.add("-border-error");
      errorIcon.classList.add("hidden");
      successIcon.classList.remove("hidden");
    } else {
      inpCon.classList.add("border-purple-700");
      inpCon.classList.remove("-border-error");
      errorIcon.classList.remove("hidden");
      successIcon.classList.add("hidden");
    }
    if (inp.value.length > 0) label.classList.add("active");
    else label.classList.remove("active");
  };
}
const persianLetters = /^[\u0600-\u06FF\s]+$/;
const numbers = /^[\u06F0-\u06F90-9]+$/;
const englishLetter = /[a-zA-Z]/;
const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
const checkInput = (val, validations, lab) => {
  let label = lab.replace("*", "");
  if (validations.required && (!val || val.length == 0))
    return `تکمیل فیلد ${label} الزامیست`;

  if (validations.persian) {
    if (!persianLetters.test(val) && val.length > 0)
      return `مقدار ${label} را به فارسی وارد کنید`;
  } else if (validations.number) {
    if (!numbers.test(val) && val.length > 0) return "فقط عدد وارد کنید";
  } else if (validations.english) {
    if (!englishLetter.test(val) && val.length > 0)
      return `مقدار ${label} را به انگلیسی وارد کنید`;
  } else if (validations.email) {
    if (!emailRegex.test(val) && val.length > 0)
      return `ایمیل خود را به درستی وارد کنید`;
  }
  if (validations.max) {
    if (val.length > validations.max)
      return `حداکثر مقدار برای این فیلد ${validations.max} می باشد`;
  }
  if (validations.length) {
    if (val.length != length) return `${label} را به درستی وارد کنید`;
  }
  if (validations.min) {
    if (val.length < validations.min)
      return `حداقل مقدار برای این فیلد ${validations.min} می باشد`;
  }

  return "";
  //   showData();
};

let allCheckForms = document.getElementsByClassName("check-form");
for (let i = 0; i < allCheckForms.length; i++) {
  let form = allCheckForms[i];
  form.addEventListener("submit", (e) => {
    e.preventDefault();
    let inputs = allCheckForms[i].getElementsByClassName("check-input");
    let errMsg = "";
    for (let j = 0; j < inputs.length; j++) {
      let inp = inputs[j].getElementsByClassName("text-input")[0];
      errMsg = checkInput(
        inp.value,
        getValidations(inp),
        inputs[j].getElementsByTagName("label")[0].textContent
      );
      let label = inputs[j].getElementsByTagName("label")[0];
      if (inp.value.length > 0) label.classList.add("active");
      else label.classList.remove("active");
    }
    if (errMsg == "") form.submit();
    else form.getElementsByClassName("form-error")[0].textContent = errMsg;
  });
}


const selects = document.getElementsByClassName("select-primary");
for (let i=0; i < selects.length; i++){
  let select = selects[i];
  let options = select.getElementsByTagName('option');
  for (let j =0; j < options.length; j++){
    options[j].setAttribute("class" , "bg-gray-900")
  }
}
