const daysSlider = document.getElementById("daysSlider");
                    const hoursSlider = document.getElementById("hoursSlider");
                    const minutesSlider = document.getElementById("minutesSlider");
                
                    const daysValue = document.getElementById("daysValue");
                    const hoursValue = document.getElementById("hoursValue");
                    const minutesValue = document.getElementById("minutesValue");
                
                    daysSlider.addEventListener("input", () => {
                      daysValue.textContent = daysSlider.value;
                    });
                
                    hoursSlider.addEventListener("input", () => {
                      hoursValue.textContent = hoursSlider.value;
                    });
                
                    minutesSlider.addEventListener("input", () => {
                      minutesValue.textContent = minutesSlider.value;
                    });