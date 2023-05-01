function getData() {
    let name = document.getElementById("name").value
    let email = document.getElementById("email").value
    let phone = document.getElementById("phone").value
    let subject = document.getElementById("subject").value
    let message = document.getElementById("message").value


if(name == ""){
    return alert("Nama Harus Diisi")
}else if(email == ""){
    return alert("Email Harus Diisi")
}else if(phone == ""){
    return alert("phone Harus Diisi")
}else if(subject == ""){
    return alert("Subject Harus Diisi")
} else if (message == ""){
    return alert("Message Harus Diisi")
}

const destination = "slametseptiawan6@gmail.com"
let a = document.createElement("a")
a.href = `mailto:${destination}?subject=${subject}&body= Hallo nama saya ${name} , saya ingin ${message}, bisakah anda menghubungi saya di ${phone}`
a.click()

let data = {
    nama: name,
    email: email,
    telp: phone,
    subject: subject,
    pesan: message,
}

console.log(data)

}



// Logic Posting Myproject


let projects = []

function getProject(event) {
    event.preventDefault()

    let pName = document.getElementById("project-name").value
    let startDate = document.getElementById("startDate").value
    let endDate = document.getElementById("endDate").value
    let description = document.getElementById("description").value
    let node = document.getElementById("nodeBox").checked
    let next = document.getElementById("nextBox").checked
    let react = document.getElementById("reactBox").checked
    let typeScript = document.getElementById("typeScriptBox").checked
    let image = document.getElementById("image").files


    image = URL.createObjectURL(image[0]);

    if (pName == "") {
        return alert("Name Project Tidak Boleh Kosong")
    } else if (startDate == "") {
        return alert("Start Tidak Boleh Kosong")
    } else if (endDate == "") {
        return alert("End Tidak Boleh Kosong")
    } else if (description == "") {
        return alert("description Tidak Boleh Kosong")
    } else if (image == "") {
        return alert("image Tidak Boleh Kosong")
    }


    let project = {
        pName,
        startDate,
        endDate,
        postAt : new Date(),
        description,
        node,
        next,
        react,
        typeScript,
        image
    }

    projects.push(project)
    console.log(projects)
    renderProject()
}


function renderProject() {
    document.getElementById("projects").innerHTML = ""
    for (let i = 0; i < projects.length; i++) {
        document.getElementById("projects").innerHTML += `
        <div class="cardBodyProject">
            <div class="img">
                <img src="${projects[i].image}" alt="">
            </div>
            <div class="cardDesripsion">
                <a id="link" href="detailProject.html">
                    <h3>${projects[i].pName}</h3>
                </a>
                <span>durasi : ${createDuration(projects[i].startDate, projects[i].endDate)}<br>${getPosted(projects[i].postAt)}</span>
                <p>${projects[i].description}</p>
            </div>
            <div class="imgtech">
                ${projects[i].next ? '<img src="asset/images/next.png" alt="">':'' }
                ${projects[i].node ? '<img src="asset/images/node.png" alt="">':'' }
                ${projects[i].react ? '<img src="asset/images/react.png" alt="">':'' }
                ${projects[i].typeScript ? '<img src="asset/images/typescript.png" alt="">':'' }     
            </div>
            <div class="control-button" style="margin: 10px;">
                <a href="">Edit</a>
                <a href="">Delete</a>
            </div>
        </div>
        `
    }
}


// Logic time

function createDuration(startDate, endDate) {
    let distance = new Date(endDate) - new Date(startDate)
    // Valid
    const dayDistance = Math.floor(distance / (24 * 60 * 60 * 1000))
    const monthDistance = Math.floor(distance / (30 * 24 * 60 * 60 * 1000))
    const yearDistance = Math.floor(distance / (12 * 30 * 24 * 60 * 60 * 1000))

    if (dayDistance > 0 && dayDistance <= 29) {
        return `${dayDistance} Hari `
    } else if (dayDistance >= 30 && monthDistance <= 12) {
        return monthDistance + " Bulan"
    } else if (monthDistance >= 12) {
        return yearDistance + " Tahun"
    } else if (dayDistance >= 0 && dayDistance <= 24) {
        return "1 hari"
    }
}

function getPosted(time) {
    const posted = new Date() - new Date(time)
    const dayPosted = Math.floor(posted / (24 * 60 * 60 * 1000))
    if (dayPosted > 0) {
        return dayPosted + " day ago"
    } else {
        const hourPosted = Math.floor(posted / (60 * 60 * 1000))
        if (hourPosted > 0) {
            return hourPosted + " hour ago"
        } else {
            const minutesPosted = Math.floor(posted / (60 * 1000))
            if (minutesPosted > 0) {
                return minutesPosted + " minutes ago"
            } else {
                const secondPosted = Math.floor(posted / 1000)
                if (secondPosted > 0) {
                    return secondPosted + " second ago"
                }
            }
        }
    }
}
setInterval(renderBlog, 1000)