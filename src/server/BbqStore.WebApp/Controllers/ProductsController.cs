using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using BbqStore.Core.Entities;
using BbqStore.Core.Services;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;

namespace BbqStore.WebApp.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class ProductsController : ControllerBase
    {
        public ProductsController(IProductService productService)
        {
            ProductService = productService;
        }

        protected IProductService ProductService { get; set; }
         
        // GET: api/Products
        [HttpGet]
        public IActionResult Get()
        {
            return Ok(ProductService.GetAll());
        }

        // GET: api/Products/5
        [HttpGet("{id:guid}", Name = "GetById")]
        public IActionResult GetById(Guid id)
        {
            return Ok(ProductService.GetById(id));
        }

        [HttpGet("{key}", Name = "GetByKey")]
        public IActionResult GetByKey(string key)
        {
            return Ok(ProductService.GetByKey(key));
        }

        // POST: api/Products
        [HttpPost]
        public IActionResult Post([FromBody] Product value)
        {
            var product = ProductService.Save(value);
            return Created($"/api/products/{product.Key}", product);
        }

        // PUT: api/Products/5
        [HttpPut("{id}")]
        public IActionResult Put(Guid id, [FromBody] Product value)
        {
            var product = ProductService.Save(value);
            return Created($"/api/products/{product.Key}", product);
        }

        // DELETE: api/ApiWithActions/5
        [HttpDelete("{id}")]
        public void Delete(int id)
        {
        }
    }
}
