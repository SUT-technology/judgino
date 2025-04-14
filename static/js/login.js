let error = {
  first_name: null,
  last_name: null,
  test: 2,
  email: 2,
};
let backdrop;
let mY = 0;
let mX = 0;
backdrop = document.getElementById("backdrop");
document.addEventListener("mousemove", function (event) {
  mX = event.clientX - 200;
  mY = event.clientY - 200;
  backdrop.style =  `--X:${mX}px; --Y:${mY}px;`;
});
const tiltEffectSettings = {
  max: 12,
  perspective: 2500,
  scale: 1.1,
  speed: 1000,
  easing: "cubic-bezier(.03,.98,.52,.99)",
};
let card;
card = document.getElementById("myCard");
document.addEventListener("mouseover", handleMouseEnter);
document.addEventListener("mouseover", handleMouseMove);
// document.addEventListener('mouseover', handleMouseLeave);
function handleMouseEnter(event) {
  setTransition(event);
}

function handleMouseMove(event) {
  const cardWidth = card.offsetWidth;
  const cardHeight = card.offsetHeight;
  const centerX = card.offsetLeft + cardWidth / 2;
  const centerY = card.offsetTop + cardHeight / 2;
  const mouseX = event.clientX - centerX;
  const mouseY = event.clientY - centerY;
  const rotateXUncapped =
    (+1 * tiltEffectSettings.max * mouseY) / (cardHeight / 2);
  const rotateYUncapped =
    (-1 * tiltEffectSettings.max * mouseX) / (cardWidth / 2);
  const rotateX =
    rotateXUncapped < -tiltEffectSettings.max
      ? -tiltEffectSettings.max
      : rotateXUncapped > tiltEffectSettings.max
      ? tiltEffectSettings.max
      : rotateXUncapped;
  const rotateY =
    rotateYUncapped < -tiltEffectSettings.max
      ? -tiltEffectSettings.max
      : rotateYUncapped > tiltEffectSettings.max
      ? tiltEffectSettings.max
      : rotateYUncapped;

  card.style.transform = `perspective(${tiltEffectSettings.perspective}px) rotateX(${rotateX}deg) rotateY(${rotateY}deg)
                          scale3d(${tiltEffectSettings.scale}, ${tiltEffectSettings.scale}, ${tiltEffectSettings.scale})`;
}

function handleMouseLeave(event) {
  console.log(event.currentTarget);
  event.currentTarget.style.transform = `perspective(${tiltEffectSettings.perspective}px) rotateX(0deg) rotateY(0deg) scale3d(1, 1, 1)`;
  setTransition(event);
}

function setTransition(event) {
  clearTimeout(card.transitionTimeoutId);
  card.style.transition = `transform ${tiltEffectSettings.speed}ms ${tiltEffectSettings.easing}`;
  card.transitionTimeoutId = setTimeout(() => {
    card.style.transition = "";
  }, tiltEffectSettings.speed);
}




const toggleForm = (el)=>{
  let loginForm = document.getElementById("login-form")
  let registerForm = document.getElementById("register-form")
  let loginContainer = document.getElementById("login-container")
  let registerContainer = document.getElementById("register-container")
  if(el.checked){
    loginContainer.classList.add("opacity-0");
    setTimeout(()=>{
      loginForm.classList.add("hidden");
      registerForm.classList.remove("hidden");
    } , 150);
    setTimeout(()=>{
      registerContainer.classList.remove("opacity-0");
    } , 300)
  }else{
    registerContainer.classList.add("opacity-0");
    setTimeout(()=>{
      registerForm.classList.add("hidden");
      loginForm.classList.remove("hidden");
    } , 150);
    setTimeout(()=>{
      loginContainer.classList.remove("opacity-0");
    } , 300)
  }
}
