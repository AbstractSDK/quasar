use anyhow::Result as AnyResult;

use crate::multitest::common::*;
use cosmwasm_std::{testing::MockApi, to_binary, Addr, CosmosMsg, MemoryStorage, WasmMsg};
use cw_multi_test::{
    App, AppBuilder, BankKeeper, DistributionKeeper, FailingModule, StakeKeeper, WasmKeeper,
};

pub type QuasarMultiHopRouterApp = App<
    BankKeeper,
    MockApi,
    MemoryStorage,
    FailingModule<Empty, Empty, Empty>,
    WasmKeeper<Empty, Empty>,
    StakeKeeper,
    DistributionKeeper,
>;

use crate::msg::{ExecuteMsg, InstantiateMsg};

#[derive(Derivative)]
#[derivative(Debug)]
pub struct QuasarVaultSuite {
    #[derivative(Debug = "ignore")]
    pub app: QuasarMultiHopRouterApp,
    // The account that deploys everything
    pub deployer: Addr,
    // executor address
    pub executor: Addr,
    // user address
    pub user: Addr,
    // router address
    pub router: Addr,
}

impl QuasarVaultSuite {
    pub fn init(init_msg: InstantiateMsg, funds: Vec<Coin>) -> Result<QuasarVaultSuite> {
        let genesis_funds = vec![coin(150000, DENOM), coin(150000, LOCAL_DENOM)];
        let deployer = Addr::unchecked(DEPLOYER);
        let executor = Addr::unchecked(EXECUTOR);
        let user = Addr::unchecked(USER);
        let mut app = AppBuilder::new().build(|router, _, storage| {
            router
                .bank
                .init_balance(storage, &deployer, genesis_funds)
                .unwrap();
        });
        app.send_tokens(
            deployer.clone(),
            user.clone(),
            &[coin(50000, DENOM), coin(50000, LOCAL_DENOM)],
        )?;
        app.send_tokens(
            deployer.clone(),
            executor.clone(),
            &[coin(50000, DENOM), coin(50000, LOCAL_DENOM)],
        )?;

        let router_id = app.store_code(contract());

        let addr = app.instantiate_contract(
            router_id,
            deployer.clone(),
            &init_msg,
            &funds,
            "router-contract",
            Some(deployer.clone().to_string()),
        )?;
        Ok(QuasarVaultSuite {
            app,
            deployer,
            executor,
            user,
            router: addr,
        })
    }

    pub fn execute(
        &mut self,
        sender: Addr,
        msg: ExecuteMsg,
        funds: Vec<Coin>,
    ) -> AnyResult<AppResponse> {
        self.app.execute(
            sender,
            CosmosMsg::Wasm(WasmMsg::Execute {
                contract_addr: self.router.to_string(),
                msg: to_binary(&msg)?,
                funds,
            }),
        )
    }
}
