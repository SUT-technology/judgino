document
  .querySelectorAll('.accordion-custom input[type="checkbox"]')
  .forEach(function (checkbox) {
    checkbox.addEventListener("change", function () {
      const accordionBody = this.nextElementSibling.nextElementSibling;
      const arrow = this.nextElementSibling.getElementsByClassName("arrow")[0];
      console.log(arrow)
      if (this.checked) {
        arrow.style.transform = "rotate(180deg)";
        accordionBody.style.maxHeight = accordionBody.scrollHeight + "px";
      } else {
        arrow.style.transform = "rotate(0deg)";
        accordionBody.style.maxHeight = 0;
      }
    });
  });
