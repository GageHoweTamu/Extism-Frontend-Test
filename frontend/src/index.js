import createPlugin from '@extism/extism';

// Create a custom function that encodes your WASM file
const getWasmModule = async () => {
  const wasmModule = await import('./plugin.wasm');
  return wasmModule.default;
};

async function runPlugin() {
  const wasmUrl = await getWasmModule();
  
  // For Extism, you might need to convert the URL to bytes
  const response = await fetch(wasmUrl);
  const wasmBytes = await response.arrayBuffer();
  
  const plugin = await createPlugin(
    new Uint8Array(wasmBytes),
    { useWasi: true }
  );
 
  const input = "Hello World";
  let out = await plugin.call("count_vowels", input);
  console.log(out.text());
}
runPlugin();