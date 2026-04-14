export async function myfunction(element,name) {
  const value = await new Promise((resolve)=>{
    setTimeout(()=>{resolve(`Greetings: ${name}`)},500);
  });
  element.dispatchEvent( new CustomEvent('mycustomevent', {detail:{value}}) );
}

globalThis.myfunction=myfunction;

class MyComponent extends HTMLElement {
  static get observedAttributes() {
    return ['src'];
  }

  constructor() { 
    super(); 
    const template=document.getElementById("gerbau").content;
    const shadowRoot=this.attachShadow({mode: "open"});
    shadowRoot.appendChild(document.importNode(template,true));
  }

  attributeChangedCallback(name, oldValue, newValue) {
    const value = `You entered: ${newValue}`;
    this.dispatchEvent(
      new CustomEvent('mycustomevent', {detail: {value}})
    );
  }
}

customElements.define('my-component', MyComponent);
