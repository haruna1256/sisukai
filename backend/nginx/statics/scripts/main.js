const status_h1 = document.getElementById('status_h1');

async function Refresh() {
    const req = await fetch('/app/refresh', {
        method: 'POST',
    });

    if (req.status !== 200) {
        //200番以外
        return false;
    }

    //更新確定
    const sreq = await fetch('/app/refreshs', {
        method: 'POST',
    })

    if (sreq.status !== 200) {
        //200番以外
        return false;
    }

    return true;
}

Refresh().then((result) => {
    if (!result) {
        status_h1.textContent = "ログインしていません";
        return;
    }

    status_h1.textContent = "ログイン中です";
});

const logout_btn = document.getElementById('logout_btn');

logout_btn.addEventListener('click', () => {
    fetch('/app/logout', {
        method: 'POST',
    }).then((res) => {
        if (res.status === 200) {
            status_h1.textContent = "ログアウトしました";
        }
    })
})

const get_user_btn = document.getElementById("get_user_btn");

async function get_user() {
    const req = await fetch('/app/getuser', {
        method: 'GET',
    });

    if (req.status !== 200) {
        //200番以外
        return false;
    }

    console.log(await req.json());
};

get_user_btn.addEventListener('click', () => {
    get_user();
})