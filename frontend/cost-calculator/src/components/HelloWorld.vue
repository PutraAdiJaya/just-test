<template>
  <div class="container mx-auto px-4">
    <h1 class="text-2xl font-bold mb-4">Perhitungan Biaya Produksi untuk tiap Kemasan Snaki</h1>
    <div class="grid grid-cols-2 gap-4"> 
      <div class="card">
        <h2 class="text-lg font-bold mb-2">Bahan yang dibutuhkan :</h2>
        <div class="mb-4">
          <label for="gula" class="block mb-2">Jumlah Gula (gram):</label>
          <input id="gula" type="text" v-model="snaki.sugar" class="w-full py-1 px-2 border border-gray-300 rounded text-white text-center" style="background-color: #000000;" />
        </div>
        <div class="mb-4">
          <label for="tapioka" class="block mb-2">Jumlah Tepung Tapioka (gram):</label>
          <input id="tapioka" type="text" v-model="snaki.tapioca" class="w-full py-1 px-2 border border-gray-300 rounded text-white text-center" style="background-color: #000000;" />
        </div>
        <div class="mb-4">
          <label for="coklat" class="block mb-2">Jumlah Coklat Padat (gram):</label>
          <input id="coklat" type="text" v-model="snaki.chocolate" class="w-full py-1 px-2 border border-gray-300 rounded text-white text-center" style="background-color: #000000;" />
        </div>
        <div class="mb-4">
          <label for="kemasan" class="block mb-2">Jumlah Kemasan:</label>
          <input id="kemasan" type="text" v-model="snaki.packaging" class="w-full py-1 px-2 border border-gray-300 rounded text-white text-center" style="background-color: #000000;" />
        </div>
      </div>
       
      <div class="card">
        <h2 class="text-lg font-bold mb-2">Harga Bahan : </h2>
        <div class="mb-4">
          <label for="gula_harga" class="block mb-2">Harga Gula per Kg:</label>
          <input id="gula_harga" type="text" v-model="cost.sugar" class="w-full py-1 px-2 border border-gray-300 rounded text-white text-center" style="background-color: #000000;" />
        </div>
        <div class="mb-4">
          <label for="tapioka_harga" class="block mb-2">Harga Tepung Tapioka per Kg:</label>
          <input id="tapioka_harga" type="text" v-model="cost.tapioca" class="w-full py-1 px-2 border border-gray-300 rounded text-white text-center" style="background-color: #000000;" />
        </div>
        <div class="mb-4">
          <label for="coklat_harga" class="block mb-2">Harga Coklat Padat per Kg:</label>
          <input id="coklat_harga" type="text" v-model="cost.chocolate" class="w-full py-1 px-2 border border-gray-300 rounded text-white text-center" style="background-color: #000000;" />
        </div>
        <div class="mb-4">
          <label for="kemasan_harga" class="block mb-2">Harga Kemasan per biji:</label>
          <input id="kemasan_harga" type="text" v-model="cost.packaging" class="w-full py-1 px-2 border border-gray-300 rounded text-white text-center" style="background-color: #000000;" />
        </div>
      </div>
    </div>
    
    <button @click="sendRequest" class="bg-blue-500 w-full  hover:bg-blue-700 text-white py-2 px-4 rounded mt-4">Hitung</button>
    
    <div v-if="result" class="mt-4">
      <div class="card">  
        <h2 class="text-lg font-bold mb-2">Biaya Produksi per Kemasan: Rp {{ result }}</h2>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      snaki: {
        sugar: 30,
        tapioca: 120,
        chocolate: 40,
        packaging: 1
      },
      cost: {
        sugar: 10000,
        tapioca: 25000,
        chocolate: 35000,
        packaging: 500
      },
      result: null
    };
  },
  methods: {
    sendRequest() { 
      const data = {
        snaki: this.snaki,
        cost: this.cost
      };
 
      fetch('http://localhost:8080/production-cost', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
      })
        .then(response => response.json())
        .then(data => { 
          this.result = data.production_cost;
        })
        .catch(error => {
          console.error('Terjadi kesalahan:', error);
        });
    }
  }
};
</script>

<style scoped>
.card {
  background-color: #ffffff;
  border-radius: 0.5rem;
  padding: 1.5rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  color: black;
}
</style>
