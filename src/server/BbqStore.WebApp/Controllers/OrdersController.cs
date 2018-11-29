using System;
using BbqStore.Core.Entities;
using BbqStore.Core.Services;
using Microsoft.AspNetCore.Mvc;

namespace BbqStore.WebApp.Controllers
{
    [ApiController]
    [Route("api/[controller]")]
    public class OrdersController : ControllerBase
    {
        public OrdersController(IOrderService orderService, IProductService productService)
        {
            OrderService = orderService;
            ProductService = productService;
        }

        protected IOrderService OrderService { get; set; }
        protected IProductService ProductService { get; set; }

        public IActionResult Get()
        {
            return Ok();
        }

        [HttpGet("{id:guid}")]
        public IActionResult GetById(Guid id)
        {
            var order = OrderService.GetById(id);
            return Ok(CreateDisplayOrder(order));
        }



        [HttpGet("{id:guid}/display")]
        public IActionResult GetForDisplay(Guid id)
        {
            var order = OrderService.GetById(id);
            var displayOrder = CreateDisplayOrder(order);

            return Ok(displayOrder);
        }

        private DisplayOrder CreateDisplayOrder(Order order)
        {
            var displayOrder = new DisplayOrder
            {
                Id = order.Id,
                CustomerName = order.CustomerName,
                Status = order.Status,
                CreatedBy = order.CreatedBy,
                CreatedDate = order.CreatedDate,
                ModifiedBy = order.ModifiedBy,
                ModifiedDate = order.ModifiedDate
            };

            order.Lines.ForEach(o =>
            {
                var product = ProductService.GetById(o.ProductId);
                var dlo = new DisplayOrderLine(o);

                if (dlo != null)
                {
                    dlo.ProductName = product.Name;
                    dlo.ProductUnit = product.Unit;

                    displayOrder.Lines.Add(dlo);
                }
            });
            return displayOrder;
        }

        [HttpPost]
        public IActionResult Post([FromBody] Order value)
        {
            var order = OrderService.Save(value);
            return Created($"/api/orders/{order.Id}", CreateDisplayOrder(order));
        }

        // PUT: api/Products/5
        [HttpPut("{id}")]
        public IActionResult Put(Guid id, [FromBody] Order value)
        {
            var order = OrderService.Save(value);
            return Accepted($"/api/orders/{order.Id}", CreateDisplayOrder(order));
        }

        [HttpPost("{id}/items")]
        public IActionResult AddToCart(Guid id, [FromBody] OrderLine line)
        {
            if (Guid.Empty.Equals(line.Id))
                line.Id = Guid.NewGuid();

            var order = OrderService.AddItemToOrder(id, line);
            return Accepted($"/api/orders/{order.Id}", CreateDisplayOrder(order));
        }

        // DELETE: api/ApiWithActions/5
        [HttpDelete("{id}")]
        public void Delete(Guid id)
        {
        }
    }
}