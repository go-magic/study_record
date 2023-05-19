var window = {}
function loadWebAssembly(bytes) {
    instance = WebAssembly.instantiate(bytes, {})
    window.q = instance.exports.encode
   return function (){
        var    t1 = parseInt(Date.parse(new Date())/1000/2);
        var    t2 = parseInt(Date.parse(new Date())/1000/2 - Math.floor(Math.random() * (50) + 1));
        return window.q(t1, t2).toString() + '|' + t1 + '|' + t2;
    }
}