

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract LostItems {


    struct LostItem {
        uint256 LostItemID;
        address  User; // 原始报告失物者的ID
        address  FoundOne; // 找到失物者的ID，如果失物未被找到则为0
        string LostItemName;
        string Description;
        uint256 LostDate;
        string Location;
        string ItemCategory;
        bool IsFound;
        uint256 FoundDate;
        uint256 reward;
        mapping(uint256 => FoundApplication) FoundApplications;
        uint256 FoundApplicationsCount;

    }

    mapping(uint256 => LostItem) public lostItems;
    uint256 public nextLostItemID;
    uint256 public TotalLostItems;

    event LostItemAdded(uint256 lostItemID, address user, string lostItemName, string location);
    function addLostItem(
        string memory _lostItemName,
        string memory _description,
        uint256 _lostDate,
        string memory _location,
        string memory _itemCategory,
        uint256 _reward
    ) public {
        // 生成一个新的失物ID
        uint256 _lostItemID = nextLostItemID++;

        // 创建失物条目
        LostItem storage newLostItem = lostItems[_lostItemID];
        newLostItem.LostItemID = _lostItemID;
        newLostItem.User = payable(msg.sender);
        newLostItem.FoundOne = payable(address(0));
        newLostItem.LostItemName = _lostItemName;
        newLostItem.Description = _description;
        newLostItem.LostDate = _lostDate;
        newLostItem.Location = _location;
        newLostItem.ItemCategory = _itemCategory;
        newLostItem.IsFound = false;
        newLostItem.FoundDate = 0;
        newLostItem.reward = _reward;
        newLostItem.FoundApplicationsCount = 0;

        TotalLostItems++;


        addHistoryRecord(_lostItemID,"AddLostItem",block.timestamp);

    }

    function updateLostItem(
        uint256 _lostItemID,
        string memory _lostItemName,
        string memory _description,
        uint256  _lostDate,
        string memory _location,
        string memory _itemCategory,
        bool _isFound,
        uint256  _foundDate,
        uint256 _reward
    ) public {
        // 检查失物ID是否存在
        require(_lostItemID >= 0 && lostItems[_lostItemID].LostItemID == _lostItemID, "Lost item ID does not exist.");
        require(msg.sender == lostItems[_lostItemID].User,"account address should be the original one");

        // 更新失物信息
        lostItems[_lostItemID].LostItemName = _lostItemName;
        lostItems[_lostItemID].Description = _description;
        lostItems[_lostItemID].LostDate = _lostDate;
        lostItems[_lostItemID].Location = _location;
        lostItems[_lostItemID].ItemCategory = _itemCategory;
        lostItems[_lostItemID].IsFound = _isFound;
        lostItems[_lostItemID].FoundDate = _foundDate;
        lostItems[_lostItemID].reward = _reward;

        addHistoryRecord(_lostItemID, "updateLostItem", block.timestamp);

    }

    function getLostItemOwner(uint256 _lostItemID) public view returns (address) {
        return lostItems[_lostItemID].User;
    }

    struct FoundApplication {
        uint256 applicationId;
        uint256 relateItemID;
        address mysender;
        address receiver;
        uint256 submitDate;
        bool result;
        bool isConfirmed;
        bool rewardPaid;
    }
    function ApplicationInfo(uint256 _lostItemID, uint256 _applicationId) public view returns (FoundApplication memory) {
        LostItem storage item = lostItems[_lostItemID];
        return item.FoundApplications[_applicationId];
    }

    uint256 public nextApplicationId =0;
    function submitFoundApplication(uint256 _lostItemID) public {
        LostItem storage item = lostItems[_lostItemID];
        require(item.User != msg.sender, "You cannot submit an application for your own lost item.");
        require(!item.IsFound, "The lost item has already been found.");

        uint256 applicationId = nextApplicationId++;
        FoundApplication storage newApplication = item.FoundApplications[applicationId];
        newApplication.applicationId = applicationId;
        newApplication.relateItemID = _lostItemID;
        newApplication.mysender = msg.sender;
        newApplication.receiver = item.User;
        newApplication.submitDate = block.timestamp;
        newApplication.isConfirmed = false;

        item.FoundApplicationsCount ++;
        // 提交申请人增加积分
        increaseUserPoints(msg.sender, 1);
    }

    function confirmFoundApplication(uint256 _lostItemID, uint256 _applicationId, bool _result) public payable {
        LostItem storage item = lostItems[_lostItemID];
        FoundApplication storage foundApp = item.FoundApplications[_applicationId];
        require(item.User == msg.sender, "Only the owner of the lost item can confirm applications.");
        require(!foundApp.isConfirmed, "Application has already been confirmed.");

        foundApp.isConfirmed = true;
        foundApp.result = _result;
        if (foundApp.result) {
            // 自动拒绝其他申请
            for (uint256 i = 0; i < item.FoundApplicationsCount; i++) {
                if (i != _applicationId && item.FoundApplications[i].relateItemID == _lostItemID) {
                    item.FoundApplications[i].isConfirmed = true;
                    item.FoundApplications[i].result = false;
                    item.FoundApplications[i].rewardPaid =false;
                    decereaseUserPoints(item.FoundApplications[i].mysender, 1);

                }
            }

            if (item.reward > 0) {
                require(address(this).balance >= item.reward, "Insufficient balance");
                (bool sent,) = payable(item.FoundOne).call{value: item.reward}("");
                require(sent, "Failed to send reward");
                foundApp.rewardPaid = true;
            }
            item.FoundOne = foundApp.mysender;
            item.IsFound = true;
            item.FoundDate = block.timestamp;
            foundApp.result = true;
            addHistoryRecord(_lostItemID,"DoneFound", block.timestamp);

            // 通过申请，申请人加5分
            increaseUserPoints(foundApp.mysender, 5);
        } else {
            foundApp.result = false;
            decereaseUserPoints(foundApp.mysender,1);
            addHistoryRecord(_lostItemID,"FailedDoneFound", block.timestamp);
        }
    }



    struct History {
        uint HistoryID;
        address Operator;
        string Operation;
        uint256 OperationTime;
        uint RelatedLostItemID;
    }

    // 存储历史记录的映射
    mapping(uint256 => History) public histories;
    // 下一个可用的历史记录ID
    uint256 public nextHistoryID;
    uint256 public TotalHistory;
    // 添加新的历史记录
    function addHistoryRecord (
        uint256 _LostItemID,
        string memory _Operation,
        uint256  _OperationTime
    ) public {
        uint256 historyID = nextHistoryID++;
        histories[historyID] = History(
            historyID,
            msg.sender, // 记录操作的用户ID
            _Operation,
            _OperationTime,
            _LostItemID // 与此历史记录相关的失物ID
        );
        TotalHistory ++;
    }

    // 使用映射存储用户的积分
    mapping(address => uint256) public userPoints;

    // 增加用户的积分
    function increaseUserPoints(address _user, uint256 _points) public {
        userPoints[_user] += _points;
    }
    function decereaseUserPoints(address _user, uint256 _points) public {
        userPoints[_user] -= _points;
    }

    // 校验申请是否存在
    function applicationExists(uint256 _lostItemID, uint256 _applicationId) public view returns (bool) {
        LostItem storage item = lostItems[_lostItemID];
        return item.FoundApplications[_applicationId].submitDate != 0;
    }

    // 校验失物是否存在
    function lostItemExists(uint256 _lostItemID) public view returns (bool) {
        return _lostItemID < nextLostItemID;
    }
    // 从智能合约中删除失物
    function deleteLostItem(uint256 _lostItemID) public {
        require(lostItemExists(_lostItemID), "Lost item does not exist.");
        require(msg.sender == lostItems[_lostItemID].User, "Only the owner can delete the lost item");
        delete lostItems[_lostItemID];
        TotalLostItems--;
    }

    // 从智能合约中删除申请
    function deleteApplication(uint256 _lostItemID, uint256 _applicationId) public {
        require(lostItemExists(_lostItemID), "Lost item does not exist.");
        require(applicationExists(_lostItemID, _applicationId), "Application does not exist.");
        delete lostItems[_lostItemID].FoundApplications[_applicationId];
        lostItems[_lostItemID].FoundApplicationsCount--;
    }


    // 管理员账户地址
    address private adminAddress ;
    // 标记是否已设置管理员地址
    bool private isAdminSet = false;

    // 构造函数，用于在部署时初始化管理员地址
    constructor() {
        setAdminAddress(msg.sender);
    }


    // 函数：设置管理员账户地址
    // _newAdminAddress：新的管理员地址
    function setAdminAddress(address _newAdminAddress) public {

        // 确保新地址不是零地址
        require(_newAdminAddress != address(0), "New admin address cannot be zero");

        // 设置管理员地址
        adminAddress = _newAdminAddress;
        // 标记管理员地址已设置
        isAdminSet = true;
    }

    // 函数：获取当前管理员地址
    function getAdminAddress() public view returns (address) {
        return adminAddress;
    }

    // 函数：向调用者转账
    function sendToCaller(uint256 amount) external payable{
        // 确保管理员地址已设置
        require(isAdminSet, "Admin address is not set");
        // 确保转账金额大于0
        require(amount > 0, "Amount must be greater than 0");
        // 确保管理员账户有足够的余额
        require(address(this).balance >= amount, "Insufficient balance");

        // 转账给调用者（msg.sender是当前调用函数的账户）
        (bool success,) = msg.sender.call{value: amount}("");
        require(success, "Transfer failed");
    }



}
