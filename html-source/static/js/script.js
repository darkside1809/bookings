let attention = Prompt();

(function () {
  'use strict'

  // Fetch all the forms we want to apply custom Bootstrap validation styles to
  let forms = document.querySelectorAll('.needs-validation')

  // Loop over them and prevent submission
  Array.prototype.slice.call(forms)
    .forEach(function (form) {
      form.addEventListener('submit', function (event) {
        if (!form.checkValidity()) {
          event.preventDefault()
          event.stopPropagation()
        }

        form.classList.add('was-validated')
      }, false)
    })
})()

document.getElementById("colorBtn").addEventListener("click", function() {
  let html = `
    <form id="check-availability-form" action="" method="post" novalidate class="needs-validation">
      <div class="row">
        <div class="col">
          <div class="row" id="reservation-dates-modal">
            <div class="col">
              <input disabled required class="form-control" type="text" name="start" id="start" placeholder="Arrival">
            </div>
            <div class="col">
              <input disabled required class="form-control" type="text" name="end" id="end" placeholder="Departure">
            </div>
          </div>
        </div>
      </div>
    </form>
  `
  attention.custom({msg: html, title: "Choose your dates"});
})

const elem = document.getElementById('reservation-dates');
const rangepicker = new DateRangePicker(elem, {
  format: "yyyy-mm-dd",
}); 

function notify(msg, msgType) {
  notie.alert({
    type: "msgType",
    text: msg,
  })
}

function notifyModal(title, text, icon, configBtn) {
  Swal.fire({
    title: title,
    html: text,
    icon: icon,
    confirmButtonText: configBtn
  })
}

function Prompt() {
    let toast = function(c) {
    const {
      msg = "",
      icon = "success",
      position = "top-end"
    } = c;

    const Toast = Swal.mixin({
      toast: true,
      title: msg, 
      position: position,
      icon: icon,
      showConfirmButton: false,
      timer: 3000,
      timerProgressBar: true,
      didOpen: (toast) => {
        toast.addEventListener('mouseenter', Swal.stopTimer)
        toast.addEventListener('mouseleave', Swal.resumeTimer)
   }
  })
    Toast.fire({})
    }

    let success = function(c) {
      const {
        msg = "",
        title = "",
        footer = "",
      } = c;
      Swal.fire({
        icon: 'success',
        title: title,
        text: msg,
        footer: footer,
      })
    }

    let error = function(c) {
      const {
        msg = "",
        title = "",
        footer = "",
      } = c;
      Swal.fire({
        icon: 'error',
        title: title,
        text: msg,
        footer: footer,
      })
    }

    async function custom(c) {
      const {
        msg = "",
        title = "",
      } = c;

      const {value: formValues } = await Swal.fire({
        title: title,
        html: msg,
        backdrop: false,
        focusConfirm: false,
        showCancelButton: true,
        willOpen: () => {
          const elem = document.getElementById('reservation-dates-modal');
          const rp = new DateRangePicker(elem, {
            format: 'yyyy-mm-dd',
            showOnFocus: true,
          })
        },
        preConfirm: () => {
          return [
            document.getElementById('start').value,
            document.getElementById('end').value
          ]
        },
        didOpen: () => {
          document.getElementById('start').removeAttribute('disabled');
          document.getElementById('end').removeAttribute('disabled');
        }
      })

      if (formValues) {
        Swal.fire(JSON.stringify(formValues))
      }
    }

    return {
      toast: toast,
      success: success,
      error: error,
      custom: custom,
    }
}